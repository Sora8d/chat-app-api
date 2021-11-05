package app

import (
	"fmt"
	"net"

	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"github.com/flydevs/chat-app-api/messaging-api/src/controllers"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/db"
	"github.com/flydevs/chat-app-api/messaging-api/src/repository/users_client"
	"github.com/flydevs/chat-app-api/messaging-api/src/services"
	"google.golang.org/grpc"
)

var (
	usersService services.MessagingService
	accroles     = map[string][]string{"/UsersProtoInterface/GetUserByUuid": {"admin"}}
)

// usersOauthService

func StartApp() {
	usersService = services.NewMessagingService(db.GetMessagingDBRepository(), users_client.GetUsersProtoClient())
	userServer := controllers.GetMessagingController(usersService)

	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMessagingProtoInterfaceServer(grpcServer, userServer)
	grpcServer.Serve(conn)
}
