package services

import (
	"context"
	"time"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/client"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/entity"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/users"
	"github.com/flydevs/chat-app-api/oauth-api/src/repository"
)

type oauthService struct {
	jwtrepo   repository.JwtRepositoryInterface
	loginrepo repository.LoginRepositoryInterface
}

type OauthServiceInterface interface {
	GenerateUser(request users.LoginRequest) (*string, *string, server_message.Svr_message)
	GenerateService(client.ServiceKey) (*string, server_message.Svr_message)
	Verify(string) (*entity.Entity, server_message.Svr_message)
}

func NewOauthService(jwtrepo repository.JwtRepositoryInterface, loginrepo repository.LoginRepositoryInterface) OauthServiceInterface {
	return &oauthService{jwtrepo: jwtrepo, loginrepo: loginrepo}
}

func (oauthsvs oauthService) GenerateUser(request users.LoginRequest) (*string, *string, server_message.Svr_message) {
	login_proto_request := request.Poblate(true, nil)
	uuid, aErr := oauthsvs.loginrepo.LoginUser(context.Background(), login_proto_request)
	if aErr != nil {
		return nil, nil, aErr
	}
	result, aErr := oauthsvs.jwtrepo.GenerateUser(*uuid)
	if aErr != nil {
		return nil, nil, aErr
	}
	return uuid, result, nil

}

func (oauthsvs oauthService) GenerateService(request client.ServiceKey) (*string, server_message.Svr_message) {
	return oauthsvs.jwtrepo.GenerateService(request)
}

func (oauthsvs oauthService) Verify(token string) (*entity.Entity, server_message.Svr_message) {
	claims, aErr := oauthsvs.jwtrepo.Verify(token)
	if aErr != nil {
		return nil, aErr
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, GetUnauthorizedErr()
	}
	return claims, nil
}

func GetUnauthorizedErr() server_message.Svr_message {
	return server_message.NewCustomMessage(401, "unauthorized token")
}
