package services

import (
	"context"
	"time"

	"github.com/Sora8d/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/client"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/entity"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/refresh_token"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/users"
	"github.com/flydevs/chat-app-api/oauth-api/src/repository"
)

type oauthService struct {
	jwtrepo         repository.JwtRepositoryInterface
	loginrepo       repository.LoginRepositoryInterface
	dbrepo          repository.DbRepositoryInterface
	tokenDuration   time.Duration
	refreshDuration time.Duration
}

type OauthServiceInterface interface {
	GenerateUser(request users.LoginRequest) (*string, []string, server_message.Svr_message)
	GenerateService(client.ServiceKey) (*string, server_message.Svr_message)
	Verify(string) (*entity.Entity, server_message.Svr_message)
	ValidateRefreshToken(token string) ([]string, server_message.Svr_message)
	BlockFamiliesByUser(uuid string) server_message.Svr_message
}

func NewOauthService(jwtrepo repository.JwtRepositoryInterface, loginrepo repository.LoginRepositoryInterface, dbrepo repository.DbRepositoryInterface, tokenDuration time.Duration, refreshDuration time.Duration) OauthServiceInterface {
	return &oauthService{jwtrepo: jwtrepo, loginrepo: loginrepo, dbrepo: dbrepo, tokenDuration: tokenDuration, refreshDuration: refreshDuration}
}

func (oauthsvs oauthService) GenerateUser(request users.LoginRequest) (*string, []string, server_message.Svr_message) {
	login_proto_request := request.Poblate(true, nil)
	uuid, aErr := oauthsvs.loginrepo.LoginUser(context.Background(), login_proto_request)
	if aErr != nil {
		return nil, nil, aErr
	}
	tokens := []string{}

	access_token_object := entity.NewEntity(time.Now().UTC().Add(oauthsvs.tokenDuration).Unix(), *uuid, 0)
	access_token, aErr := oauthsvs.jwtrepo.GenerateAccessToken(access_token_object)
	if aErr != nil {
		return nil, nil, aErr
	}
	tokens = append(tokens, *access_token)

	refreshToken_object := refresh_token.NewRefreshToken(time.Now().UTC().Add(oauthsvs.refreshDuration).Unix(), 0, *uuid, 0)
	aErr = oauthsvs.dbrepo.AddNewTokenFamily(&refreshToken_object)
	if aErr != nil {
		return nil, nil, aErr
	}
	refreshToken, aErr := oauthsvs.jwtrepo.GenerateRefreshToken(refreshToken_object)
	if aErr != nil {
		return nil, nil, aErr
	}
	tokens = append(tokens, *refreshToken)
	return uuid, tokens, nil

}

func (oauthsvs oauthService) GenerateService(request client.ServiceKey) (*string, server_message.Svr_message) {
	permissions, aErr := request.ValidateKey()
	if aErr != nil {
		return nil, aErr
	}
	access_token_object := entity.NewEntity(time.Now().UTC().Add(oauthsvs.tokenDuration).Unix(), "00000000-0000-0000-0000-000000000000", *permissions)
	access_token, aErr := oauthsvs.jwtrepo.GenerateAccessToken(access_token_object)
	if aErr != nil {
		return nil, aErr
	}
	return access_token, nil
}

func (oauthsvs oauthService) Verify(token string) (*entity.Entity, server_message.Svr_message) {
	claims, aErr := oauthsvs.jwtrepo.VerifyAccessToken(token)
	if aErr != nil {
		return nil, aErr
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, GetUnauthorizedErr()
	}
	return claims, nil
}

func (oauthsvs oauthService) ValidateRefreshToken(token string) ([]string, server_message.Svr_message) {
	refreshToken_object, aErr := oauthsvs.jwtrepo.VerifyRefreshToken(token)
	if aErr != nil {
		return nil, aErr
	}
	if refreshToken_object.ExpiresAt < time.Now().UTC().Unix() {
		return nil, GetUnauthorizedErr()
	}

	db_refreshToken, blocked, aErr := oauthsvs.dbrepo.CheckToken(refreshToken_object)
	if aErr != nil {
		return nil, aErr
	}
	if blocked {
		return nil, GetUnauthorizedErr()
	}
	if !refreshToken_object.Equal(*db_refreshToken) {
		if aErr := oauthsvs.dbrepo.BlockTokenFamily(refreshToken_object.Family); aErr != nil {
			return nil, aErr
		}
		return nil, GetUnauthorizedErr()
	}
	refreshToken_object.SetExpiration(time.Now().UTC().Add(oauthsvs.refreshDuration).Unix())
	tokens := []string{}

	access_token_object := entity.NewEntity(time.Now().UTC().Add(oauthsvs.tokenDuration).Unix(), refreshToken_object.Uuid, 0)
	access_token, aErr := oauthsvs.jwtrepo.GenerateAccessToken(access_token_object)
	if aErr != nil {
		return nil, aErr
	}
	tokens = append(tokens, *access_token)
	refreshToken, aErr := oauthsvs.jwtrepo.GenerateRefreshToken(*refreshToken_object)
	if aErr != nil {
		return nil, aErr
	}
	aErr = oauthsvs.dbrepo.AddNewToken(refreshToken_object)
	if aErr != nil {
		return nil, aErr
	}
	tokens = append(tokens, *refreshToken)
	return tokens, nil
}

func (oauthsvs oauthService) BlockFamiliesByUser(uuid string) server_message.Svr_message {
	return oauthsvs.dbrepo.BlockFamiliesByUser(uuid)
}

func GetUnauthorizedErr() server_message.Svr_message {
	return server_message.NewCustomMessage(401, "unauthorized token")
}
