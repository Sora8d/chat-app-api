package app

import (
	"fmt"
	"net"

	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
	"github.com/flydevs/chat-app-api/users-api/src/config"
	"github.com/flydevs/chat-app-api/users-api/src/controllers"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
	"github.com/flydevs/chat-app-api/users-api/src/services"
	"google.golang.org/grpc"
)

var usersService services.UsersServiceInterface

func StartApp() {
	postgresql.DbInit()
	usersService = services.NewUsersService(db.GetUserDbRepository())
	userServer := controllers.GetNewUserServer(usersService)

	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersProtoInterfaceServer(grpcServer, userServer)
	grpcServer.Serve(conn)
}
