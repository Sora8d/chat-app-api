package app

import (
	"fmt"
	"net"

	"github.com/Sora8d/common/logger"
	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/users-api/src/config"
	"github.com/flydevs/chat-app-api/users-api/src/controllers"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
	"github.com/flydevs/chat-app-api/users-api/src/repository/oauth"
	"github.com/flydevs/chat-app-api/users-api/src/server"
	"github.com/flydevs/chat-app-api/users-api/src/services"
	"google.golang.org/grpc"
)

var (
	accroles = map[string]int{
		"/flydev_chat_app_users.UsersProtoInterface/GetUserByUuid":    2,
		"/flydev_chat_app_users.UsersProtoInterface/UpdateUser":       0,
		"/flydev_chat_app_users.UsersProtoInterface/DeleteUserByUuid": 0,
		"/flydev_chat_app_users.UsersProtoInterface/SearchContact":    0,
	}
)

// usersOauthService

func StartApp() {
	postgresql.DbInit()
	userServer := server.GetNewUserServer(controllers.GetNewUserController(services.NewUsersService(db.GetUserDbRepository())))
	logger.Info(fmt.Sprint("starting up the users api at port ", config.Config["PORT"]))
	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	oauth_interceptor := services.NewAuthInterceptor(accroles, oauth.GetOauthRepository())
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(oauth_interceptor.Unary()))
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUsersProtoInterfaceServer(grpcServer, userServer)
	grpcServer.Serve(conn)
}
