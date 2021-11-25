package oauth

import (
	"context"

	"github.com/flydevs/chat-app-api/users-api/src/clients/proto_clients"
	"github.com/flydevs/chat-app-api/users-api/src/clients/rpc/oauth"
)

type oauthRepository struct{}

type OauthRepositoryInterface interface {
	LoginService()
	Verify(string) (*oauth.EntityResponse, error)
}

func GetOauthRepository() OauthRepositoryInterface {
	return &oauthRepository{}
}

func (oauthRepository) LoginService()

func (oauthRepository) Verify(jwt string) (*oauth.EntityResponse, error) {
	client := proto_clients.GetOauthClient()
	proto_jwt := oauth.JWT{Jwt: jwt}
	return client.Client.Verify(context.Background(), &proto_jwt)
}
