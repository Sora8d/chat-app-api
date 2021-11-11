package controllers

import (
	"context"

	"github.com/flydevs/chat-app-api/common/logger"
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
	CreateConversation(*proto_messaging.Conversation) (*proto_messaging.Uuid, *proto_messaging.SvrMsg)
	CreateMessage(*proto_messaging.CreateMessageRequest) (*proto_messaging.Uuid, *proto_messaging.SvrMsg)
	CreateUserConversation(*proto_messaging.CreateUserConversationRequest) *proto_messaging.SvrMsg
	GetConversationsByUser(*proto_messaging.Uuid) ([]*proto_messaging.ConversationAndParticipants, *proto_messaging.SvrMsg)
	GetMessagesByConversation(*proto_messaging.MessageRequest) ([]*proto_messaging.Message, *proto_messaging.SvrMsg)
	UpdateConversationInfo(*proto_messaging.Conversation) (*proto_messaging.Conversation, *proto_messaging.SvrMsg)
	UpdateMessage(*proto_messaging.Message) (*proto_messaging.Message, *proto_messaging.SvrMsg)
}

func GetMessagingController(messaging_service services.MessagingService) messagingController {
	return messagingController{svc: messaging_service}
}

func (mc messagingController) CreateConversation(ctx context.Context, pbc *proto_messaging.Conversation) (*proto_messaging.UuidMsg, error) {
	logger.Info("accesed this")
	var new_conversation conversation.Conversation
	new_conversation.Poblate(false, pbc)
	result_uuid, resp_msg := mc.svc.CreateConversation(new_conversation)
	var Pb_response proto_messaging.UuidMsg
	Pb_response.Msg = poblateMsg(resp_msg)
	if result_uuid != nil {
		pb_resp_uuid := proto_messaging.Uuid{}
		result_uuid.Poblate(true, &pb_resp_uuid)
		Pb_response.Uuid = &pb_resp_uuid
	}
	return &Pb_response, nil
}

func (mc messagingController) CreateMessage(ctx context.Context, pbm *proto_messaging.CreateMessageRequest) (*proto_messaging.UuidMsg, error) {
	var err server_message.Svr_message
	var conversation_uuid *proto_messaging.Uuid
	for !pbm.ConversationExists {
		data := conversation.ConversationAndParticipants{}
		data.Poblate(false, pbm.NewConvo)
		convo_uuid, msg := mc.svc.CreateConversation(data.Conversation)
		if msg.GetStatus() != 200 {
			err = msg
			break
		}
		ucs := conversation.CreateUserConversationRequest{}
		ucs.Ucs = data.Participants
		ucs.Conversation.Uuid = convo_uuid.Uuid
		msg = mc.svc.CreateUserConversation(ucs)
		if msg.GetStatus() != 200 {
			err = msg
			break
		}
		pb_convo_uuid := proto_messaging.Uuid{}
		convo_uuid.Poblate(true, &pb_convo_uuid)
		conversation_uuid = &pb_convo_uuid
		break
	}
	var Pb_response proto_messaging.UuidMsg
	if err != nil {
		Pb_response.Msg = poblateMsg(err)
		return &Pb_response, nil
	}

	var new_message message.Message
	new_message.Poblate(false, pbm.Message)
	if conversation_uuid != nil {
		new_message.ConversationUuid = conversation_uuid.Uuid
	}

	result_conversation_uuid, resp_msg := mc.svc.CreateMessage(new_message)
	Pb_response.Msg = poblateMsg(resp_msg)
	if result_conversation_uuid != nil {
		pb_resp_uuid := proto_messaging.Uuid{}
		result_conversation_uuid.Poblate(true, &pb_resp_uuid)
		Pb_response.Uuid = &pb_resp_uuid
	}
	return &Pb_response, nil
}

func (mc messagingController) CreateUserConversation(ctx context.Context, pbuc *proto_messaging.CreateUserConversationRequest) (*proto_messaging.UuidMsg, error) {
	var new_user_conversation conversation.CreateUserConversationRequest
	new_user_conversation.Poblate(pbuc.UserConversations)
	new_user_conversation.Conversation.Uuid = pbuc.ConversationUuid.Uuid
	resp_msg := mc.svc.CreateUserConversation(new_user_conversation)

	var Pb_response proto_messaging.UuidMsg
	Pb_response.Msg = poblateMsg(resp_msg)
	return &Pb_response, nil
}

//

func (mc messagingController) GetConversationsByUser(ctx context.Context, pbuuid *proto_messaging.Uuid) (*proto_messaging.ArrayConversationResponse, error) {
	result, response_msg := mc.svc.GetConversationsByUser(pbuuid.Uuid)
	var Pb_response proto_messaging.ArrayConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*proto_messaging.ConversationAndParticipants
		for _, content := range result {
			var new_pb_convo proto_messaging.ConversationAndParticipants
			content.Poblate(true, &new_pb_convo)
			pb_result = append(pb_result, &new_pb_convo)
		}
		Pb_response.Conversations = pb_result
	}
	return &Pb_response, nil
}

//

func (mc messagingController) GetMessagesByConversation(ctx context.Context, pbuuid *proto_messaging.MessageRequest) (*proto_messaging.ArrayMessageResponse, error) {
	result, response_msg := mc.svc.GetMessagesByConversation(pbuuid.UcUuid.Uuid, pbuuid.ConversationUuid.Uuid)
	var Pb_response proto_messaging.ArrayMessageResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*proto_messaging.Message
		for _, content := range result {
			var new_pb_message proto_messaging.Message
			content.Poblate(true, &new_pb_message)
			pb_result = append(pb_result, &new_pb_message)
		}
		Pb_response.Message = pb_result
	}
	return &Pb_response, nil
}

func (mc messagingController) UpdateConversationInfo(ctx context.Context, pb_convo *proto_messaging.Conversation) (*proto_messaging.UpdateConversationResponse, error) {
	var request_convo conversation.Conversation
	request_convo.Poblate(false, pb_convo)
	result, response_msg := mc.svc.UpdateConversationInfo(request_convo.Uuid, request_convo.ConversationInfo)
	var Pb_response proto_messaging.UpdateConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result proto_messaging.Conversation
		result.Poblate(true, &pb_result)
		Pb_response.Conversation = &pb_result
	}
	return &Pb_response, nil
}

func (mc messagingController) UpdateMessage(ctx context.Context, pb_message *proto_messaging.Message) (*proto_messaging.MessageMsgResponse, error) {
	result, response_msg := mc.svc.UpdateMessage(pb_message.Uuid.Uuid, pb_message.Text)
	var Pb_response proto_messaging.MessageMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result proto_messaging.Message
		result.Poblate(true, &pb_result)
		Pb_response.Message = &pb_result
	}
	return &Pb_response, nil
}

func poblateMsg(msg server_message.Svr_message) *proto_messaging.SvrMsg {
	pb_new_msg := proto_messaging.SvrMsg{
		Status:  int64(msg.GetStatus()),
		Message: msg.GetMessage(),
	}
	return &pb_new_msg
}
