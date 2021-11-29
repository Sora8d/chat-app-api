package repository

import (
	"errors"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	"github.com/dgrijalva/jwt-go"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/entity"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/refresh_token"
)

type jwtRepository struct {
	secretKey string
}

type JwtRepositoryInterface interface {
	GenerateAccessToken(claims entity.Entity) (*string, server_message.Svr_message)
	GenerateRefreshToken(refreshToken refresh_token.RefreshToken) (*string, server_message.Svr_message)
	VerifyAccessToken(string) (*entity.Entity, server_message.Svr_message)
	VerifyRefreshToken(accessToken string) (*refresh_token.RefreshToken, server_message.Svr_message)
}

func NewjwtRepository(secretKey string) JwtRepositoryInterface {
	return &jwtRepository{secretKey: secretKey}
}

func (jR jwtRepository) GenerateAccessToken(claims entity.Entity) (*string, server_message.Svr_message) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(jR.secretKey))
	if err != nil {
		logger.Error("error generating JWT in GenerateUser", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (jR jwtRepository) GenerateRefreshToken(refreshToken refresh_token.RefreshToken) (*string, server_message.Svr_message) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
	result, err := token.SignedString([]byte(jR.secretKey))
	if err != nil {
		logger.Error("error generating JWT in GenerateRefreshToken", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (jR jwtRepository) VerifyAccessToken(accessToken string) (*entity.Entity, server_message.Svr_message) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&entity.Entity{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(jR.secretKey), nil
		},
	)

	if err != nil {
		return nil, GetUnauthorizedErr()
	}

	claims, ok := token.Claims.(*entity.Entity)
	if !ok {
		return nil, GetUnauthorizedErr()
	}
	return claims, nil
}

func (jR jwtRepository) VerifyRefreshToken(refreshToken string) (*refresh_token.RefreshToken, server_message.Svr_message) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&refresh_token.RefreshToken{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(jR.secretKey), nil
		},
	)

	if err != nil {
		return nil, GetUnauthorizedErr()
	}

	claims, ok := token.Claims.(*refresh_token.RefreshToken)
	if !ok {
		return nil, GetUnauthorizedErr()
	}
	return claims, nil
}

func GetUnauthorizedErr() server_message.Svr_message {
	return server_message.NewCustomMessage(401, "unauthorized token")
}
