package app

import (
	"fmt"
	"net"
	"time"

	"github.com/Sora8d/common/logger"
	proto_oauth "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/oauth-api/src/config"
	"github.com/flydevs/chat-app-api/oauth-api/src/controller"
	"github.com/flydevs/chat-app-api/oauth-api/src/repository"
	"github.com/flydevs/chat-app-api/oauth-api/src/server"
	"github.com/flydevs/chat-app-api/oauth-api/src/services"
	"google.golang.org/grpc"
)

func StartApp() {
	var expire_duration time.Duration = 3 * time.Hour
	jwt_key := config.Config["SECRET_KEY"]

	Oauth_server := server.GetNewServer(controller.GetNewOauthController(services.NewOauthService(repository.NewjwtRepository(jwt_key, expire_duration), repository.NewLoginRepository())))
	logger.Info(fmt.Sprintf("initating app on %s...", config.Config["PORT"]))
	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto_oauth.RegisterOauthProtoInterfaceServer(grpcServer, Oauth_server)
	grpcServer.Serve(conn)
}
