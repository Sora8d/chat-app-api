package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/client"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/users"
)

type jwtRepository struct {
	secretKey     string
	tokenDuration time.Duration
}

type JwtRepositoryInterface interface {
}

func NewjwtRepository(secretKey string, tokenDuration time.Duration) JwtRepositoryInterface {
	return &jwtRepository{secretKey: secretKey, tokenDuration: tokenDuration}
}

func (jR jwtRepository) GenerateUser(user string) (*string, server_message.Svr_message) {
	claims := users.User{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(jR.tokenDuration).Unix()},
		Uuid: user,
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
	permissions, name, aErr := serviceKey.ValidateKey()
	if aErr != nil {
		return nil, aErr
	}
	claims := client.Client{StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(jR.tokenDuration).Unix()},
		Permissions: *permissions,
		ServiceName: *name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString([]byte(jR.secretKey))
	if err != nil {
		logger.Error("error generating JWT ", err)
		return nil, server_message.NewInternalError()
	}
	return &result, nil
}

func (jR jwtRepository) UserVerify(accessToken string) (*users.User, server_message.Svr_message) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&users.User{},
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

	claims, ok := token.Claims.(*users.User)
	if !ok {
		return nil, server_message.NewBadRequestError("invalid token claims")
	}
	return claims, nil
}

func (jR jwtRepository) ServiceVerify(accessToken string) (*users.User, server_message.Svr_message) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&users.User{},
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

	claims, ok := token.Claims.(*users.User)
	if !ok {
		return nil, server_message.NewBadRequestError("invalid token claims")
	}
	return claims, nil
}
