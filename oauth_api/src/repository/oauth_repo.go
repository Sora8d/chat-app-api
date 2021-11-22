package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/client"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/entity"
)

type jwtRepository struct {
	secretKey     string
	tokenDuration time.Duration
}

type JwtRepositoryInterface interface {
	GenerateUser(string) (*string, server_message.Svr_message)
	GenerateService(client.ServiceKey) (*string, server_message.Svr_message)
	UserVerify(string) (*entity.Entity, server_message.Svr_message)
	ClientVerify(string) (*entity.Entity, server_message.Svr_message)
}

func NewjwtRepository(secretKey string, tokenDuration time.Duration) JwtRepositoryInterface {
	return &jwtRepository{secretKey: secretKey, tokenDuration: tokenDuration}
}

func (jR jwtRepository) GenerateUser(user_uuid string) (*string, server_message.Svr_message) {
	claims := entity.Entity{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(jR.tokenDuration).Unix()},
		Uuid:        user_uuid,
		Permissions: 0,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString([]byte(jR.secretKey))
	if err != nil {
		logger.Error("error generating JWT ", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (jR jwtRepository) GenerateService(serviceKey client.ServiceKey) (*string, server_message.Svr_message) {
	permissions, aErr := serviceKey.ValidateKey()
	if aErr != nil {
		return nil, aErr
	}
	claims := entity.Entity{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Add(jR.tokenDuration).Unix()},
		Uuid:        "00000000-0000-0000-0000-000000000000",
		Permissions: *permissions,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString([]byte(jR.secretKey))
	if err != nil {
		logger.Error("error generating JWT ", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (jR jwtRepository) UserVerify(accessToken string) (*entity.Entity, server_message.Svr_message) {
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
		return nil, server_message.NewBadRequestError(fmt.Sprint("invalid token: ", accessToken))
	}

	claims, ok := token.Claims.(*entity.Entity)
	if !ok {
		return nil, server_message.NewBadRequestError("invalid token claims")
	}
	return claims, nil
}

func (jR jwtRepository) ClientVerify(accessToken string) (*entity.Entity, server_message.Svr_message) {
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
		return nil, server_message.NewBadRequestError(fmt.Sprint("invalid token: ", accessToken))
	}

	claims, ok := token.Claims.(*entity.Entity)
	if !ok {
		return nil, server_message.NewBadRequestError("invalid token claims")
	}
	return claims, nil
}
