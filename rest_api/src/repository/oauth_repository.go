package repository

import (
	"context"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/proto_clients"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/oauth"
)

type oauthRepository struct {
}

type OauthRepositoryInterface interface {
	LoginUser(ctx context.Context, in *oauth.LoginRequest) (*oauth.JWTAndUuidResponse, server_message.Svr_message)
}

func GetOauthRepository() OauthRepositoryInterface {
	return &oauthRepository{}
}

func (oauthRepository) LoginUser(ctx context.Context, in *oauth.LoginRequest) (*oauth.JWTAndUuidResponse, server_message.Svr_message) {
	client := proto_clients.GetOauthClient()
	response, err := client.Client.LoginUser(ctx, in)
	if err != nil {
		logger.Error("error in oauth_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response, nil
}
