package repository

import (
	"context"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/oauth-api/src/clients/proto_clients"
	proto_login "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/login"
)

type loginRepository struct {
}

type LoginRepositoryInterface interface {
	LoginUser(context.Context, *proto_login.User) (*string, server_message.Svr_message)
}

func NewLoginRepository() LoginRepositoryInterface {
	return &loginRepository{}
}

func (lr loginRepository) LoginUser(ctx context.Context, request *proto_login.User) (*string, server_message.Svr_message) {
	client := proto_clients.GetLoginClient()

	response, err := client.Client.UserLogin(ctx, request)
	if err != nil {
		logger.Error("error in LoginUser function inside login_repository", err)
		return nil, server_message.NewInternalError()
	}
	if response.Msg.Status >= 400 {
		return nil, poblateMsgfromProto(response.Msg)
	}
	return &response.Users[0].Uuid.Uuid, nil
}

func poblateMsgfromProto(pbmsg *proto_login.SvrMsg) server_message.Svr_message {
	return server_message.NewCustomMessage(int(pbmsg.Status), pbmsg.Message)
}
