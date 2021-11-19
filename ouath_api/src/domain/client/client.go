package client

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/config"
)

type Client struct {
	jwt.StandardClaims
	Permissions int
	ServiceName string
}

type ServiceKey string

func (sk ServiceKey) ValidateKey() (*int, *string, server_message.Svr_message) {
	switch string(sk) {
	case config.Config["OAUTH_KEY"]:
		perm := 2
		name := "oauth"
		return &perm, &name, nil

	case config.Config["USERS_KEY"]:
		perm := 1
		name := "users"
		return &perm, &name, nil

	case config.Config["REST_KEY"]:
		perm := 1
		name := "rest"
		return &perm, &name, nil

	case config.Config["MESSAGING_KEY"]:
		perm := 2
		name := "messaging"
		return &perm, &name, nil
	default:
		return nil, nil, server_message.NewBadRequestError("invalid credentials")
	}
}
