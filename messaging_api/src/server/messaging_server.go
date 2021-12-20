package server

import (
	"context"
	"net/http"

	"github.com/Sora8d/common/server_message"
	proto_messaging "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/controllers"
)

type messagingServer struct {
	proto_messaging.UnimplementedMessagingProtoInterfaceServer
	controller controllers.MessagingController
}

func GetMessagingserver(messaging_controller controllers.MessagingController) messagingServer {
	return messagingServer{controller: messaging_controller}
}

func (ms messagingServer) CreateConversation(ctx context.Context, pbc *proto_messaging.Conversation) (*proto_messaging.UuidMsg, error) {
	uuid, err := ms.controller.CreateConversation(ctx, pbc)
	var Proto_uuid_response proto_messaging.UuidMsg
	var proto_msg *proto_messaging.SvrMsg
	if err != nil {
		proto_msg = poblateMsg(err)
	} else {
		proto_msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "conversation created"))
	}
	Proto_uuid_response.Msg = proto_msg
	Proto_uuid_response.Uuid = uuid
	return &Proto_uuid_response, nil

}
func (ms messagingServer) CreateMessage(ctx context.Context, pbm *proto_messaging.CreateMessageRequest) (*proto_messaging.UuidMsg, error) {
	uuid, err := ms.controller.CreateMessage(ctx, pbm)
	var Proto_uuid_response proto_messaging.UuidMsg
	var proto_msg *proto_messaging.SvrMsg
	if err != nil {
		proto_msg = poblateMsg(err)
	} else {
		proto_msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "message created"))
	}
	Proto_uuid_response.Msg = proto_msg
	Proto_uuid_response.Uuid = uuid
	return &Proto_uuid_response, nil
}

func (ms messagingServer) CreateUserConversation(ctx context.Context, pbuc *proto_messaging.CreateUserConversationRequest) (*proto_messaging.SvrMsg, error) {
	err := ms.controller.CreateUserConversation(ctx, pbuc)
	var msg *proto_messaging.SvrMsg
	if err != nil {
		msg = poblateMsg(err)
	} else {
		msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "user_conversation/s created"))
	}
	return msg, nil
}

func (ms messagingServer) GetConversationsByUser(ctx context.Context, pbuuid *proto_messaging.Uuid) (*proto_messaging.ArrayConversationResponse, error) {
	proto_conversation_participants, err := ms.controller.GetConversationsByUser(ctx, pbuuid)
	var Proto_response proto_messaging.ArrayConversationResponse
	var msg *proto_messaging.SvrMsg
	if err != nil {
		msg = poblateMsg(err)
	} else {
		msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "conversations retrieved"))
	}
	Proto_response.Msg = msg
	Proto_response.Conversations = proto_conversation_participants
	return &Proto_response, nil
}

func (ms messagingServer) GetMessagesByConversation(ctx context.Context, getMessagesRequest *proto_messaging.GetMessages) (*proto_messaging.ArrayMessageResponse, error) {
	messages, err := ms.controller.GetMessagesByConversation(ctx, getMessagesRequest)
	var Proto_response proto_messaging.ArrayMessageResponse
	var msg *proto_messaging.SvrMsg
	if err != nil {
		msg = poblateMsg(err)
	} else {
		msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "messages retrieved"))
	}
	Proto_response.Msg = msg
	Proto_response.Message = messages
	return &Proto_response, nil
}

func (ms messagingServer) UpdateConversationInfo(ctx context.Context, pb_convo *proto_messaging.Conversation) (*proto_messaging.UpdateConversationResponse, error) {
	proto_conversation_updated, err := ms.controller.UpdateConversationInfo(ctx, pb_convo)
	var Proto_response proto_messaging.UpdateConversationResponse
	var msg *proto_messaging.SvrMsg
	if err != nil {
		msg = poblateMsg(err)
	} else {
		msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "conversation info updated"))
	}
	Proto_response.Msg = msg
	Proto_response.Conversation = proto_conversation_updated
	return &Proto_response, nil
}

func (ms messagingServer) UpdateMessage(ctx context.Context, pb_message *proto_messaging.Message) (*proto_messaging.MessageMsgResponse, error) {
	proto_message_updated, err := ms.controller.UpdateMessage(ctx, pb_message)
	var Proto_response proto_messaging.MessageMsgResponse
	var msg *proto_messaging.SvrMsg
	if err != nil {
		msg = poblateMsg(err)
	} else {
		msg = poblateMsg(server_message.NewCustomMessage(http.StatusOK, "message updated"))
	}
	Proto_response.Msg = msg
	Proto_response.Message = proto_message_updated
	return &Proto_response, nil
}

func poblateMsg(msg server_message.Svr_message) *proto_messaging.SvrMsg {
	pb_new_msg := proto_messaging.SvrMsg{
		Status:  int64(msg.GetStatus()),
		Message: msg.GetMessage(),
	}
	return &pb_new_msg
}
