package repository

import (
	"context"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	"github.com/flydevs/chat-app-api/rest-api/src/clients/proto_clients"
	proto_users "github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/users"
)

type usersRepository struct {
}

type UsersRepositoryInterface interface {
	CreateUser(context.Context, *proto_users.RegisterUser) server_message.Svr_message
	GetUserProfileByUuid(context.Context, *proto_users.MultipleUuids) ([]*proto_users.UserProfile, server_message.Svr_message)
	UpdateUser(context.Context, *proto_users.UpdateUserRequest) ([]*proto_users.UserProfile, server_message.Svr_message)
}

func GetUsersRepository() UsersRepositoryInterface {
	return &usersRepository{}
}

func (ur usersRepository) CreateUser(ctx context.Context, in *proto_users.RegisterUser) server_message.Svr_message {
	client := proto_clients.GetUsersClient()
	response, err := client.Client.CreateUser(ctx, in)
	if err != nil {
		logger.Error("error in users_repository,", err)
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(int(response.Status), response.Message)
}

func (ur usersRepository) GetUserProfileByUuid(ctx context.Context, in *proto_users.MultipleUuids) ([]*proto_users.UserProfile, server_message.Svr_message) {
	client := proto_clients.GetUsersClient()
	response, err := client.Client.GetUserProfileByUuid(ctx, in)
	if err != nil {
		logger.Error("error in users_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.User, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (ur usersRepository) UpdateUser(ctx context.Context, in *proto_users.UpdateUserRequest) ([]*proto_users.UserProfile, server_message.Svr_message) {
	client := proto_clients.GetUsersClient()
	response, err := client.Client.UpdateUser(ctx, in)
	if err != nil {
		logger.Error("error in users_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.User, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}
