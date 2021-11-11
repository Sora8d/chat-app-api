package server

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
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
	uuid, err := ms.controller.CreateConversation(pbc)
}
func (ms messagingServer) CreateMessage(ctx context.Context, pbm *proto_messaging.CreateMessageRequest) (*proto_messaging.UuidMsg, error) {
	uuid, err := ms.controller.CreateMessage(pbm)
}

func (ms messagingServer) CreateUserConversation(ctx context.Context, pbuc *proto_messaging.CreateUserConversationRequest) (*proto_messaging.UuidMsg, error) {
	err := ms.controller.CreateUserConversation(pbuc)
}
func (ms messagingServer) GetConversationsByUser(ctx context.Context, pbuuid *proto_messaging.Uuid) (*proto_messaging.ArrayConversationResponse, error) {
	convos, err := ms.controller.GetConversationsByUser(pbuuid)
}

func (ms messagingServer) GetMessagesByConversation(ctx context.Context, pbuuid *proto_messaging.MessageRequest) (*proto_messaging.ArrayMessageResponse, error) {
	messages, err := ms.controller.GetMessagesByConversation(pbuuid)
}

func (ms messagingServer) UpdateConversationInfo(ctx context.Context, pb_convo *proto_messaging.Conversation) (*proto_messaging.UpdateConversationResponse, error) {
	convo, err := ms.controller.UpdateConversationInfo(pb_convo)
}

func (ms messagingServer) UpdateMessage(ctx context.Context, pb_message *proto_messaging.Message) (*proto_messaging.MessageMsgResponse, error) {
	message, err := ms.controller.UpdateMessage(pb_message)
}

func poblateMsg(msg server_message.Svr_message) *proto_messaging.SvrMsg {
	pb_new_msg := proto_messaging.SvrMsg{
		Status:  int64(msg.GetStatus()),
		Message: msg.GetMessage(),
	}
	return &pb_new_msg
}
