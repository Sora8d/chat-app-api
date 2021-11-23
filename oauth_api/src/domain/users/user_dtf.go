package users

import (
	proto_users "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/login"
	proto_oauth "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/oauth"
)

func (lr *LoginRequest) Poblate(direction_out bool, obj *proto_oauth.LoginRequest) *proto_users.User {
	if direction_out {
		object_to_return := proto_users.User{}
		object_to_return.LoginUser = lr.Username
		object_to_return.LoginPassword = lr.Password
		return &object_to_return
	} else {
		lr.Username = obj.Username
		lr.Password = obj.Password
		return nil
	}
}
