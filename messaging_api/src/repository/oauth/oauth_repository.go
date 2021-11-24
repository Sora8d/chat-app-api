package oauth

import (
	"context"
	"fmt"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/messaging-api/src/clients/proto_clients"
	"github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"google.golang.org/grpc/metadata"
)

type oauthRepository struct {
	current_access_token *string
}

type OauthRepositoryInterface interface {
	LoginService() (context.Context, server_message.Svr_message)
	Verify(string) (*oauth.EntityResponse, error)
}

func GetOauthRepository() OauthRepositoryInterface {
	return &oauthRepository{}
}

func (oR *oauthRepository) LoginService() (context.Context, server_message.Svr_message) {
	client := proto_clients.GetOauthClient()
	valid := false
	if oR.current_access_token != nil {
		oauth_jwt := oauth.JWT{Jwt: *oR.current_access_token}
		response, _ := client.Client.Verify(context.Background(), &oauth_jwt)
		if response != nil && response.Response.Status == 200 {
			valid = true
		}
	}
	if !valid {
		proto_key := oauth.ServiceKey{Key: config.Config["SECRET_KEY"]}
		response, err := client.Client.LoginClient(context.Background(), &proto_key)
		if err != nil {
			logger.Error("Login service error ", err)
			return nil, server_message.NewInternalError()
		}
		if response.Response.Status != 200 {
			logger.Error("Login service error in response body", fmt.Errorf("%d, %s", response.Response.Status, response.Response.Message))
			return nil, server_message.NewInternalError()
		}
		oR.setServiceToken(&response.Jwt)
	}
	md := metadata.New(map[string]string{"access_token": *oR.current_access_token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return ctx, nil
}

func (oauthRepository) Verify(jwt string) (*oauth.EntityResponse, error) {
	client := proto_clients.GetOauthClient()
	proto_jwt := oauth.JWT{Jwt: jwt}
	return client.Client.Verify(context.Background(), &proto_jwt)
}

func (oR *oauthRepository) setServiceToken(tk *string) {
	oR.current_access_token = tk
}
