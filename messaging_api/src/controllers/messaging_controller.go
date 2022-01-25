package controllers

import (
	"context"
	"errors"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	proto_messaging "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/oauth"
	oauth_repo "github.com/flydevs/chat-app-api/messaging-api/src/repository/oauth"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
	"google.golang.org/grpc/metadata"
)

type messagingController struct {
	proto_messaging.UnimplementedMessagingProtoInterfaceServer

	svc       services.MessagingService
	oauthrepo oauth_repo.OauthRepositoryInterface
}

type MessagingController interface {
	CreateConversation(context.Context, *proto_messaging.Conversation) (*proto_messaging.Uuid, server_message.Svr_message)
	CreateMessage(context.Context, *proto_messaging.CreateMessageRequest) (*proto_messaging.Uuid, server_message.Svr_message)
	CreateUserConversation(context.Context, *proto_messaging.CreateUserConversationRequest) server_message.Svr_message

	KickUser(ctx context.Context, proto_kick_request *proto_messaging.KickUserRequest) server_message.Svr_message

	GetConversationsByUser(context.Context, *proto_messaging.Uuid) ([]*proto_messaging.ConversationAndParticipants, server_message.Svr_message)
	GetMessagesByConversation(context.Context, *proto_messaging.GetMessages) ([]*proto_messaging.Message, server_message.Svr_message)
	UpdateConversationInfo(context.Context, *proto_messaging.Conversation) (*proto_messaging.Conversation, server_message.Svr_message)
	UpdateMessage(context.Context, *proto_messaging.Message) (*proto_messaging.Message, server_message.Svr_message)
}

func GetMessagingController(messaging_service services.MessagingService, oauthrepo oauth_repo.OauthRepositoryInterface) MessagingController {
	return messagingController{svc: messaging_service, oauthrepo: oauthrepo}
}

func (mc messagingController) CreateConversation(ctx context.Context, pbc *proto_messaging.Conversation) (*proto_messaging.Uuid, server_message.Svr_message) {
	_, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}

	var new_conversation conversation.Conversation
	new_conversation.Poblate(false, pbc)
	result_uuid, aErr := mc.svc.CreateConversation(new_conversation)
	if aErr != nil {
		return nil, aErr
	}

	proto_uuid := proto_messaging.Uuid{}

	result_uuid.Poblate(true, &proto_uuid)

	return &proto_uuid, nil
}

func (mc messagingController) CreateMessage(ctx context.Context, pbm *proto_messaging.CreateMessageRequest) (*proto_messaging.Uuid, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	at_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}
	if ok := compareUuids(*at_uuid, pbm.Message.AuthorUuid.Uuid); !ok {
		return nil, oauth.GetError("2")
	}

	ctx_with_client_token, aErr := mc.oauthrepo.LoginService()
	if aErr != nil {
		return nil, aErr
	}

	var conversation_uuid proto_messaging.Uuid
	if pbm.CreateConversation {
		if pbm.NewConvo == nil {
			return nil, server_message.NewBadRequestError("no conversation data provided")
		}
		data := conversation.ConversationAndParticipants{}
		data.Poblate(false, pbm.NewConvo)
		convo_uuid, err := mc.svc.CreateConversation(data.Conversation)
		if err != nil {
			return nil, err
		}
		ucs := conversation.CreateUserConversationRequest{}
		ucs.SetUserConversationSlice(data.Participants)
		ucs.SetConversation(conversation.Conversation{Uuid: convo_uuid.Uuid})
		err = mc.svc.CreateUserConversation(ctx_with_client_token, *at_uuid, false, ucs)
		if err != nil {
			return nil, err
		}
		convo_uuid.Poblate(true, &conversation_uuid)
	}

	var new_message message.Message
	new_message.Poblate(false, pbm.Message)
	if new_message.ConversationUuid == "" {
		new_message.SetConversationUuid(conversation_uuid.Uuid)
	}

	result_conversation_uuid, err := mc.svc.CreateMessage(ctx_with_client_token, *at_uuid, new_message)
	if err != nil {
		return nil, err
	}
	pb_resp_uuid := proto_messaging.Uuid{}
	result_conversation_uuid.Poblate(true, &pb_resp_uuid)
	return &pb_resp_uuid, nil
}

func (mc messagingController) CreateUserConversation(ctx context.Context, pbuc *proto_messaging.CreateUserConversationRequest) server_message.Svr_message {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return aErr
	}
	verification_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return aErr
	}
	ctx_with_client_token, aErr := mc.oauthrepo.LoginService()
	if aErr != nil {
		return aErr
	}
	var new_user_conversation conversation.CreateUserConversationRequest
	new_user_conversation.Poblate(pbuc.UserConversations)
	new_user_conversation.SetConversation(conversation.Conversation{Uuid: pbuc.ConversationUuid.Uuid})
	err := mc.svc.CreateUserConversation(ctx_with_client_token, *verification_uuid, true, new_user_conversation)
	return err
}

//

func (mc messagingController) GetConversationsByUser(ctx context.Context, proto_user_uuid *proto_messaging.Uuid) ([]*proto_messaging.ConversationAndParticipants, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	at_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}
	if ok := compareUuids(*at_uuid, proto_user_uuid.Uuid); !ok {
		return nil, oauth.GetError("2")
	}
	conversation_participants, err := mc.svc.GetConversationsByUser(proto_user_uuid.Uuid)
	if err != nil {
		return nil, err
	}
	proto_conversation_participants := conversation_participants.Poblate(nil)
	return proto_conversation_participants, nil
}

//
func (mc messagingController) GetMessagesByConversation(ctx context.Context, convo_uuid *proto_messaging.GetMessages) ([]*proto_messaging.Message, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	verification_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}
	messages, err := mc.svc.GetMessagesByConversation(*verification_uuid, convo_uuid.Uuid.Uuid, convo_uuid.BeforeDate, convo_uuid.AfterDate)
	if err != nil {
		return nil, err
	}
	proto_messages := messages.Poblate(nil)
	return proto_messages, nil
}

func (mc messagingController) UpdateConversationInfo(ctx context.Context, pb_convo *proto_messaging.Conversation) (*proto_messaging.Conversation, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	verification_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}

	var request_convo conversation.Conversation
	request_convo.Poblate(false, pb_convo)
	conversation_updated, err := mc.svc.UpdateConversationInfo(request_convo.Uuid, *verification_uuid, request_convo.ConversationInfo)
	if err != nil {
		return nil, err
	}
	var proto_conversation_updated proto_messaging.Conversation
	conversation_updated.Poblate(true, &proto_conversation_updated)
	return &proto_conversation_updated, nil
}

func (mc messagingController) UpdateMessage(ctx context.Context, pb_message *proto_messaging.Message) (*proto_messaging.Message, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	verification_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}

	message_updated, err := mc.svc.UpdateMessage(pb_message.Uuid.Uuid, *verification_uuid, pb_message.Text)
	if err != nil {
		return nil, err
	}
	var proto_message_updated proto_messaging.Message
	message_updated.Poblate(true, &proto_message_updated)
	return &proto_message_updated, nil
}
func (mc messagingController) KickUser(ctx context.Context, proto_kick_request *proto_messaging.KickUserRequest) server_message.Svr_message {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return aErr
	}
	verification_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return aErr
	}

	return mc.svc.KickUser(proto_kick_request.UserConversation.GetUuid(), proto_kick_request.Conversation.GetUuid(), *verification_uuid)
}

func validateContext(ctx context.Context) (metadata.MD, server_message.Svr_message) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("error reading metadata", errors.New("error in validatecontext, no metadata"))
		return nil, server_message.NewInternalError()
	}
	at_err := md.Get("status")
	if at_err == nil {
		logger.Error("error reading metadata", errors.New("error in validatecontext, status is nil"))
		return nil, server_message.NewInternalError()
	}
	return md, oauth.GetError(at_err[0])
}

func fetchUuid(md metadata.MD) (*string, server_message.Svr_message) {
	at_uuid := md.Get("uuid")
	if at_uuid == nil {
		logger.Error("error reading metadata", errors.New("error in UpdateUser, uuid is nil"))
		return nil, server_message.NewInternalError()
	}
	return &at_uuid[0], nil
}

func compareUuids(uuid1, uuid2 string) bool {
	return uuid1 == uuid2
}
