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

var (
	usersService services.UsersServiceInterface
	accroles     = map[string][]string{"/UsersProtoInterface/GetUserByUuid": {"admin"}}
)

// usersOauthService

func StartApp() {
	postgresql.DbInit()
	usersService = services.NewUsersService(db.GetUserDbRepository())
	userServer := controllers.GetNewUserServer(usersService)

	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	oauth_interceptor := services.NewAuthInterceptor(accroles)
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(oauth_interceptor.Unary()))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersProtoInterfaceServer(grpcServer, userServer)
	grpcServer.Serve(conn)
}
