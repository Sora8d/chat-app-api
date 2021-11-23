package client

import (
	"github.com/Sora8d/common/server_message"
	"github.com/dgrijalva/jwt-go"
	"github.com/flydevs/chat-app-api/oauth-api/src/config"
)

type Client struct {
	jwt.StandardClaims
	Permissions int
	ServiceName string
}

type ServiceKey string

func (sk *ServiceKey) Set(str string) {
	*sk = ServiceKey(str)
}

func (sk ServiceKey) ValidateKey() (*int, server_message.Svr_message) {
	switch string(sk) {
	case config.Config["OAUTH_KEY"], config.Config["MESSAGING_KEY"]:
		perm := 2
		return &perm, nil

	case config.Config["USERS_KEY"], config.Config["REST_KEY"]:
		perm := 1
		return &perm, nil
	default:
		return nil, server_message.NewBadRequestError("invalid credentials")
	}
}
