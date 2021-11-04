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

	GetConversation(*pb.Uuid) (*pb.ConvMsgResponse, error)
	GetMessage(*pb.Uuid) (*pb.MessageMsgResponse, error)
	GetUserConversation(*pb.Uuid) (*pb.UCMsgResponse, error)

	GetConversationsByUser(*pb.Uuid) (*pb.ArrayConversationResponse, error)
	GetMessagesByAuthor(*pb.Uuid) (*pb.ArrayMessageResponse, error)
	GetUserConversationForUser(*pb.Uuid) (*pb.ArrayUserConversationResponse, error)

	GetMessagesByConversation(*pb.Uuid) (*pb.ArrayMessageResponse, error)
	GetUserConversationForConversation(*pb.Uuid) (*pb.ArrayUserConversationResponse, error)

	UpdateConversationLastMsg(*pb.UuidandMessageUuid) (*pb.ConvMsgResponse, error)
	UpdateUserConversation(*pb.UuidandMessageUuid) (*pb.UCMsgResponse, error)

	UpdateConversationInfo(*pb.Conversation) (*pb.ConvMsgResponse, error)
	UpdateMessage(*pb.Message) (*pb.MessageMsgResponse, error)
}

func GetMessagingControler(messaging_service services.MessagingService) MessagingController {
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

func (mc *messagingController) GetConversation(pbuuid *pb.Uuid) (*pb.ConvMsgResponse, error) {
	result, response_msg := mc.svc.GetConversationByUuid(pbuuid.Uuid)
	var Pb_response pb.ConvMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.Conversation
		result.Poblate(true, &pb_result)
		Pb_response.Convo = &pb_result
	}
	return &Pb_response, nil
}

func (mc *messagingController) GetMessage(pbuuid *pb.Uuid) (*pb.MessageMsgResponse, error) {
	result, response_msg := mc.svc.GetMessageByUuid(pbuuid.Uuid)
	var Pb_response pb.MessageMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.Message
		result.Poblate(true, &pb_result)
		Pb_response.Message = &pb_result
	}
	return &Pb_response, nil
}

func (mc *messagingController) GetUserConversation(pbuuid *pb.Uuid) (*pb.UCMsgResponse, error) {
	result, response_msg := mc.svc.GetUserConversationByUuid(pbuuid.Uuid)
	var Pb_response pb.UCMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.UserConversation
		result.Poblate(true, &pb_result)
		Pb_response.UserConversation = &pb_result
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) GetConversationsByUser(pbuuid *pb.Uuid) (*pb.ArrayConversationResponse, error) {
	result, response_msg := mc.svc.GetConversationsByUser(pbuuid.Uuid)
	var Pb_response pb.ArrayConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*pb.Conversation
		for _, content := range result {
			var new_pb_convo pb.Conversation
			content.Poblate(true, &new_pb_convo)
			pb_result = append(pb_result, &new_pb_convo)
		}
		Pb_response.Conversation = pb_result
	}
	return &Pb_response, nil
}

func (mc *messagingController) GetMessagesByAuthor(pbuuid *pb.Uuid) (*pb.ArrayMessageResponse, error) {
	result, response_msg := mc.svc.GetMessagesByAuthor(pbuuid.Uuid)
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
func (mc *messagingController) GetUserConversationForUser(pbuuid *pb.Uuid) (*pb.ArrayUserConversationResponse, error) {
	result, response_msg := mc.svc.GetUserConversationsForUser(pbuuid.Uuid)
	var Pb_response pb.ArrayUserConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*pb.UserConversation
		for _, content := range result {
			var new_pb_uc pb.UserConversation
			content.Poblate(true, &new_pb_uc)
			pb_result = append(pb_result, &new_pb_uc)
		}
		Pb_response.UserConversation = pb_result
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) GetMessagesByConversation(pbuuid *pb.Uuid) (*pb.ArrayMessageResponse, error) {
	result, response_msg := mc.svc.GetMessagesByConversation(pbuuid.Uuid)
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
func (mc *messagingController) GetUserConversationForConversation(pbuuid *pb.Uuid) (*pb.ArrayUserConversationResponse, error) {
	result, response_msg := mc.svc.GetUserConversationsForConversation(pbuuid.Uuid)
	var Pb_response pb.ArrayUserConversationResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result []*pb.UserConversation
		for _, content := range result {
			var new_pb_uc pb.UserConversation
			content.Poblate(true, &new_pb_uc)
			pb_result = append(pb_result, &new_pb_uc)
		}
		Pb_response.UserConversation = pb_result
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) UpdateConversationLastMsg(pbuuids *pb.UuidandMessageUuid) (*pb.ConvMsgResponse, error) {
	result, response_msg := mc.svc.UpdateConversationLastMsg(pbuuids.OgEntity.Uuid, pbuuids.Message.Uuid)
	var Pb_response pb.ConvMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)

	if result != nil {
		var pb_result pb.Conversation
		result.Poblate(true, &pb_result)
		Pb_response.Convo = &pb_result
	}
	return &Pb_response, nil
}
func (mc *messagingController) UpdateUserConversation(pbuuids *pb.UuidandMessageUuid) (*pb.UCMsgResponse, error) {
	result, response_msg := mc.svc.UserConversationUpdateLastAccess(pbuuids.OgEntity.Uuid, pbuuids.Message.Uuid)
	var Pb_response pb.UCMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)

	if result != nil {
		var pb_result pb.UserConversation
		result.Poblate(true, &pb_result)
		Pb_response.UserConversation = &pb_result
	}
	return &Pb_response, nil
}

//

func (mc *messagingController) UpdateConversationInfo(pb_convo *pb.Conversation) (*pb.ConvMsgResponse, error) {
	var request_convo conversation.Conversation
	request_convo.Poblate(false, pb_convo)
	result, response_msg := mc.svc.UpdateConversationInfo(request_convo.Uuid, request_convo.ConversationInfo)
	var Pb_response pb.ConvMsgResponse
	Pb_response.Msg = poblateMsg(response_msg)
	if result != nil {
		var pb_result pb.Conversation
		result.Poblate(true, &pb_result)
		Pb_response.Convo = &pb_result
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
