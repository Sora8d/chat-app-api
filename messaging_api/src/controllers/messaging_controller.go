package controllers

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
	proto_messaging "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
)

type messagingController struct {
	proto_messaging.UnimplementedMessagingProtoInterfaceServer

	svc services.MessagingService
}

type MessagingController interface {
	CreateConversation(context.Context, *proto_messaging.Conversation) (*proto_messaging.Uuid, server_message.Svr_message)
	CreateMessage(context.Context, *proto_messaging.CreateMessageRequest) (*proto_messaging.Uuid, server_message.Svr_message)
	CreateUserConversation(context.Context, *proto_messaging.CreateUserConversationRequest) server_message.Svr_message
	GetConversationsByUser(context.Context, *proto_messaging.Uuid) ([]*proto_messaging.ConversationAndParticipants, server_message.Svr_message)
	GetMessagesByConversation(context.Context, *proto_messaging.MessageRequest) ([]*proto_messaging.Message, server_message.Svr_message)
	UpdateConversationInfo(context.Context, *proto_messaging.Conversation) (*proto_messaging.Conversation, server_message.Svr_message)
	UpdateMessage(context.Context, *proto_messaging.Message) (*proto_messaging.Message, server_message.Svr_message)
}

func GetMessagingController(messaging_service services.MessagingService) MessagingController {
	return messagingController{svc: messaging_service}
}

func (mc messagingController) CreateConversation(ctx context.Context, pbc *proto_messaging.Conversation) (*proto_messaging.Uuid, server_message.Svr_message) {
	var new_conversation conversation.Conversation
	new_conversation.Poblate(false, pbc)
	result_uuid, err := mc.svc.CreateConversation(new_conversation)
	proto_uuid := proto_messaging.Uuid{}
	if result_uuid != nil {
		result_uuid.Poblate(true, &proto_uuid)
	}
	return &proto_uuid, err
}

func (mc messagingController) CreateMessage(ctx context.Context, pbm *proto_messaging.CreateMessageRequest) (*proto_messaging.Uuid, server_message.Svr_message) {
	var conversation_uuid proto_messaging.Uuid
	if !pbm.ConversationExists {
		data := conversation.ConversationAndParticipants{}
		data.Poblate(false, pbm.NewConvo)
		convo_uuid, err := mc.svc.CreateConversation(data.Conversation)
		if err != nil {
			return nil, err
		}
		ucs := conversation.CreateUserConversationRequest{}
		ucs.Ucs = data.Participants
		ucs.Conversation.Uuid = convo_uuid.Uuid
		err = mc.svc.CreateUserConversation(ucs)
		if err != nil {
			return nil, err
		}
		convo_uuid.Poblate(true, &conversation_uuid)
	}

	var new_message message.Message
	new_message.Poblate(false, pbm.Message)
	if new_message.ConversationUuid == "" {
		new_message.ConversationUuid = conversation_uuid.Uuid
	}

	result_conversation_uuid, err := mc.svc.CreateMessage(new_message)
	if err != nil {
		return nil, err
	}
	pb_resp_uuid := proto_messaging.Uuid{}
	result_conversation_uuid.Poblate(true, &pb_resp_uuid)
	return &pb_resp_uuid, nil
}

func (mc messagingController) CreateUserConversation(ctx context.Context, pbuc *proto_messaging.CreateUserConversationRequest) server_message.Svr_message {
	var new_user_conversation conversation.CreateUserConversationRequest
	new_user_conversation.Poblate(pbuc.UserConversations)
	new_user_conversation.Conversation.Uuid = pbuc.ConversationUuid.Uuid
	err := mc.svc.CreateUserConversation(new_user_conversation)
	return err
}

//

func (mc messagingController) GetConversationsByUser(ctx context.Context, proto_user_uuid *proto_messaging.Uuid) ([]*proto_messaging.ConversationAndParticipants, server_message.Svr_message) {
	conversation_participants, err := mc.svc.GetConversationsByUser(proto_user_uuid.Uuid)
	if err != nil {
		return nil, err
	}
	var proto_conversation_participants []*proto_messaging.ConversationAndParticipants
	for _, content := range conversation_participants {
		var new_pb_convo proto_messaging.ConversationAndParticipants
		content.Poblate(true, &new_pb_convo)
		proto_conversation_participants = append(proto_conversation_participants, &new_pb_convo)
	}
	return proto_conversation_participants, nil
}

//

func (mc messagingController) GetMessagesByConversation(ctx context.Context, pbuuid *proto_messaging.MessageRequest) ([]*proto_messaging.Message, server_message.Svr_message) {
	messages, err := mc.svc.GetMessagesByConversation(pbuuid.UcUuid.Uuid, pbuuid.ConversationUuid.Uuid)
	if err != nil {
		return nil, err
	}
	var proto_messages []*proto_messaging.Message
	for _, content := range messages {
		var new_pb_message proto_messaging.Message
		content.Poblate(true, &new_pb_message)
		proto_messages = append(proto_messages, &new_pb_message)
	}
	return proto_messages, nil
}

func (mc messagingController) UpdateConversationInfo(ctx context.Context, pb_convo *proto_messaging.Conversation) (*proto_messaging.Conversation, server_message.Svr_message) {
	var request_convo conversation.Conversation
	request_convo.Poblate(false, pb_convo)
	conversation_updated, err := mc.svc.UpdateConversationInfo(request_convo.Uuid, request_convo.ConversationInfo)
	if err != nil {
		return nil, err
	}
	var proto_conversation_updated proto_messaging.Conversation
	conversation_updated.Poblate(true, &proto_conversation_updated)
	return &proto_conversation_updated, nil
}

func (mc messagingController) UpdateMessage(ctx context.Context, pb_message *proto_messaging.Message) (*proto_messaging.Message, server_message.Svr_message) {
	message_updated, err := mc.svc.UpdateMessage(pb_message.Uuid.Uuid, pb_message.Text)
	if err != nil {
		return nil, err
	}
	var proto_message_updated proto_messaging.Message
	message_updated.Poblate(true, &proto_message_updated)
	return &proto_message_updated, nil
}
