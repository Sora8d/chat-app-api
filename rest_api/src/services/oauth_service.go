package services

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/oauth"
	"github.com/flydevs/chat-app-api/rest-api/src/domain"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
)

type oauthService struct {
	oauthrepo repository.OauthRepositoryInterface
}

type OauthServiceInterface interface {
	LoginUser(*oauth.LoginRequest) (domain.Response, *string)
}

func NewOauthService(oauthrepo repository.OauthRepositoryInterface) OauthServiceInterface {
	return &oauthService{oauthrepo: oauthrepo}
}

func (oauthsvs oauthService) LoginUser(request *oauth.LoginRequest) (domain.Response, *string) {
	ctx := context.Background()
	response, err := oauthsvs.oauthrepo.LoginUser(ctx, request)
	if err != nil {
		return Response.CreateResponse(nil, err), nil
	}
	return Response.CreateResponse(response.Uuid, server_message.NewCustomMessage(int(response.Response.Status), response.Response.Message)), &response.Jwt
}
