package app

import (
	"fmt"
	"net"

	"github.com/Sora8d/common/logger"
	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
	"github.com/flydevs/chat-app-api/users-api/src/config"
	"github.com/flydevs/chat-app-api/users-api/src/controllers"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
	"github.com/flydevs/chat-app-api/users-api/src/server"
	"github.com/flydevs/chat-app-api/users-api/src/services"
	"google.golang.org/grpc"
)

var (
	accroles = map[string][]string{"/UsersProtoInterface/GetUserByUuid": {"admin"}}
)

// usersOauthService

func StartApp() {
	postgresql.DbInit()
	userServer := server.GetNewUserServer(controllers.GetNewUserController(services.NewUsersService(db.GetUserDbRepository())))
	logger.Info(fmt.Sprintln("initiating app on port ", config.Config["PORT"]))
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
