package app

import (
	"fmt"
	"net"

	"github.com/Sora8d/common/logger"
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"github.com/flydevs/chat-app-api/messaging-api/src/controllers"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/oauth"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/twilio"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
	"github.com/flydevs/chat-app-api/messaging-api/src/server"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
	"google.golang.org/grpc"
)

var (
	messagingService    services.MessagingService
	messagingController controllers.MessagingController

	accroles = map[string]int{
		"/flydevs_chat_app_messaging.MessagingProtoInterface/CreateConversation":        0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/GetConversationsByUser":    0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateConversationInfo":    0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/CreateMessage":             0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/GetMessagesByConversation": 0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/UpdateMessage":             0,
		"/flydevs_chat_app_messaging.MessagingProtoInterface/CreateUserConversation":    0,
	}
)

func StartApp() {
	messagingService = services.NewMessagingService(db.GetMessagingDBRepository(), users_client.GetUsersProtoClient(), twilio.GetTwilioMock())
<<<<<<< HEAD
	messagingController = controllers.GetMessagingController(messagingService)
=======
	messagingController = controllers.GetMessagingController(messagingService, oauth.GetOauthRepository())
>>>>>>> main
	messagingServer := server.GetMessagingserver(messagingController)
	logger.Info(fmt.Sprintf("initating app on %s...", config.Config["PORT"]))
	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	oauth_interceptor := services.NewAuthInterceptor(accroles, oauth.GetOauthRepository())
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(oauth_interceptor.Unary()))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMessagingProtoInterfaceServer(grpcServer, messagingServer)
	grpcServer.Serve(conn)
}
