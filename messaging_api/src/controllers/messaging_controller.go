package controllers

import (
	"github.com/flydevs/chat-app-api/common/server_message"
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/conversation"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/message"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
)

type messagingController struct {
	pb.UnimplementedMessagingProtoInterfaceServer

	svc services.MessagingService
}

type MessagingController interface {
	CreateConversation(*pb.Conversation) (*pb.UuidMsg, error)
	CreateMessage(*pb.Message) (*pb.UuidMsg, error)
	CreateUserConversation(*pb.UserConversation) (*pb.UuidMsg, error)
	//Later this will be done with the JWT information instead of a provided uuid.
	GetConversationsByUser(*pb.Uuid) (*pb.ArrayConversationResponse, error)
	GetMessagesByConversation(*pb.MessageRequest) (*pb.ArrayMessageResponse, error)
	//
	UpdateConversationInfo(*pb.Conversation) (*pb.UpdateConversationResponse, error)
	UpdateMessage(*pb.Message) (*pb.MessageMsgResponse, error)
}

func GetMessagingController(messaging_service services.MessagingService) MessagingController {
	return &messagingController{svc: messaging_service}
}

func (mc *messagingController) CreateConversation(pbc *pb.Conversation) (*pb.UuidMsg, error) {
	var new_conversation conversation.Conversation
	new_conversation.Poblate(false, pbc)
	result_uuid, resp_msg := mc.svc.CreateConversation(new_conversation)
	var Pb_response pb.UuidMsg
	Pb_response.Msg = poblateMsg(resp_msg)
	if result_uuid != nil {
		pb_resp_uuid := pb.Uuid{}
		result_uuid.Poblate(true, &pb_resp_uuid)
		Pb_response.Uuid = &pb_resp_uuid
	}
	return &Pb_response, nil
}

func (mc *messagingController) CreateMessage(pbm *pb.Message) (*pb.UuidMsg, error) {
	var new_message message.Message
	new_message.Poblate(false, pbm)
	result_uuid, resp_msg := mc.svc.CreateMessage(new_message)

	var Pb_response pb.UuidMsg
	Pb_response.Msg = poblateMsg(resp_msg)
	if result_uuid != nil {
		pb_resp_uuid := pb.Uuid{}
		result_uuid.Poblate(true, &pb_resp_uuid)
		Pb_response.Uuid = &pb_resp_uuid
	}
	return &Pb_response, nil
}

func (mc *messagingController) CreateUserConversation(pbuc *pb.UserConversation) (*pb.UuidMsg, error) {
	var new_user_conversation conversation.UserConversation
	new_user_conversation.Poblate(false, pbuc)
	result_uuid, resp_msg := mc.svc.CreateUserConversation(new_user_conversation)

	var Pb_response pb.UuidMsg
	Pb_response.Msg = poblateMsg(resp_msg)
	if result_uuid != nil {
		pb_resp_uuid := pb.Uuid{}
		result_uuid.Poblate(true, &pb_resp_uuid)
		Pb_response.Uuid = &pb_resp_uuid
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) GetConversationsByUser(pbuuid *pb.Uuid) (*pb.ArrayConversationResponse, error) {
	result, response_msg := mc.svc.GetConversationsByUser(pbuuid.Uuid)
	var Pb_response pb.ArrayConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*pb.ConversationAndParticipants
		for _, content := range result {
			var new_pb_convo pb.ConversationAndParticipants
			content.Poblate(true, &new_pb_convo)
			pb_result = append(pb_result, &new_pb_convo)
		}
		Pb_response.Conversations = pb_result
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) GetMessagesByConversation(pbuuid *pb.MessageRequest) (*pb.ArrayMessageResponse, error) {
	result, response_msg := mc.svc.GetMessagesByConversation(pbuuid.UserUuid.Uuid, pbuuid.ConversationUuid.Uuid)
	var Pb_response pb.ArrayMessageResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*pb.Message
		for _, content := range result {
			var new_pb_message pb.Message
			content.Poblate(true, &new_pb_message)
			pb_result = append(pb_result, &new_pb_message)
		}
		Pb_response.Message = pb_result
	}
	return &Pb_response, nil
}

func (mc *messagingController) UpdateConversationInfo(pb_convo *pb.Conversation) (*pb.UpdateConversationResponse, error) {
	var request_convo conversation.Conversation
	request_convo.Poblate(false, pb_convo)
	result, response_msg := mc.svc.UpdateConversationInfo(request_convo.Uuid, request_convo.ConversationInfo)
	var Pb_response pb.UpdateConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.Conversation
		result.Poblate(true, &pb_result)
		Pb_response.Conversation = &pb_result
	}
	return &Pb_response, nil
}

func (mc *messagingController) UpdateMessage(pb_message *pb.Message) (*pb.MessageMsgResponse, error) {
	result, response_msg := mc.svc.UpdateMessage(pb_message.Uuid.Uuid, pb_message.Text)
	var Pb_response pb.MessageMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.Message
		result.Poblate(true, &pb_result)
		Pb_response.Message = &pb_result
	}
	return &Pb_response, nil
}

func poblateMsg(msg server_message.Svr_message) *pb.SvrMsg {
	pb_new_msg := pb.SvrMsg{
		Status:  int64(msg.GetStatus()),
		Message: msg.GetMessage(),
	}
	return &pb_new_msg
}
