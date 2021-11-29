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
	LoginUser(ctx context.Context, in *oauth.LoginRequest) (*oauth.JWTwRrefreshUuidResponse, server_message.Svr_message)
	ValidateRefreshToken(ctx context.Context, in *oauth.JWT) (*oauth.JWTwRrefreshUuidResponse, server_message.Svr_message)
	RevokeUsersTokens(ctx context.Context, in *oauth.Uuid) server_message.Svr_message
}

func GetOauthRepository() OauthRepositoryInterface {
	return &oauthRepository{}
}

func (oauthRepository) LoginUser(ctx context.Context, in *oauth.LoginRequest) (*oauth.JWTwRrefreshUuidResponse, server_message.Svr_message) {
	client := proto_clients.GetOauthClient()
	response, err := client.Client.LoginUser(ctx, in)
	if err != nil {
		logger.Error("error in oauth_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response, nil
}

func (oauthRepository) ValidateRefreshToken(ctx context.Context, in *oauth.JWT) (*oauth.JWTwRrefreshUuidResponse, server_message.Svr_message) {
	client := proto_clients.GetOauthClient()
	response, err := client.Client.ValidateRefreshToken(ctx, in)
	if err != nil {
		logger.Error("error in validate_refresh_token,", err)
		return nil, server_message.NewInternalError()
	}
	return response, nil
}

func (oauthRepository) RevokeUsersTokens(ctx context.Context, in *oauth.Uuid) server_message.Svr_message {
	client := proto_clients.GetOauthClient()
	response, err := client.Client.RevokeUsersTokens(ctx, in)
	if err != nil {
		logger.Error("error in revoke_users_tokens,", err)
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(int(response.GetStatus()), response.GetMessage())
}
