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

var (
	dbrepo repository.DbRepositoryInterface
)

func StartApp() {
	var expire_duration time.Duration = time.Duration(config.Config["ACCESSTOKEN_EXPIRATION"].(int)) * time.Minute
	var refresh_token_duration time.Duration = time.Duration(config.Config["REFRESH_EXPIRATION"].(int)) * time.Minute
	jwt_key := config.Config["SECRET_KEY"].(string)

	dbrepo = repository.NewDbRepository()

	Oauth_server := server.GetNewServer(controller.GetNewOauthController(
		services.NewOauthService(
			repository.NewjwtRepository(jwt_key),
			repository.NewLoginRepository(),
			dbrepo,
			expire_duration, refresh_token_duration)))
	logger.Info(fmt.Sprintf("initating app on %s...", config.Config["PORT"]))
	conn, err := net.Listen("tcp", config.Config["PORT"].(string))
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto_oauth.RegisterOauthProtoInterfaceServer(grpcServer, Oauth_server)
	go cleanTokens()
	grpcServer.Serve(conn)

}
func cleanTokens() {
	for {
		fmt.Println("running")
		expiration_int := config.Config["REFRESH_EXPIRATION"].(int)
		time.Sleep(time.Duration(expiration_int) * time.Minute)
		intervals := fmt.Sprintf("%d minutes", expiration_int)
		err := dbrepo.CleanTokens(intervals)
		if err != nil {
			logger.Error("error cloaning tokens", err)
		}
	}
}
