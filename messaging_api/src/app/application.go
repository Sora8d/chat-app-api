package app

import (
	"fmt"
	"net"

	"github.com/Sora8d/common/logger"
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"github.com/flydevs/chat-app-api/messaging-api/src/controllers"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/twilio"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
	"github.com/flydevs/chat-app-api/messaging-api/src/server"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
	"google.golang.org/grpc"
)

var (
	messagingService    services.MessagingService
	messagingController controllers.MessagingController

//	accroles         = map[string][]string{"/UsersProtoInterface/GetUserByUuid": {"admin"}}
)

// usersOauthService

func StartApp() {
	messagingService = services.NewMessagingService(db.GetMessagingDBRepository(), users_client.GetUsersProtoClient(), twilio.GetTwilioMock())
	messagingController = controllers.GetMessagingController(messagingService)
	messagingServer := server.GetMessagingserver(messagingController)
	logger.Info(fmt.Sprintf("initating app on %s...", config.Config["PORT"]))
	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMessagingProtoInterfaceServer(grpcServer, messagingServer)
	grpcServer.Serve(conn)
}
