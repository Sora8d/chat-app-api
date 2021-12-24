package services

import (
	"context"

	"github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/rest-api/src/domain"
	"github.com/flydevs/chat-app-api/rest-api/src/repository"
)

type usersService struct {
	users_repo repository.UsersRepositoryInterface
}

type UsersServiceInterface interface {
	CreateUser(context.Context, *users.RegisterUser) domain.Response
	GetUserProfileByUuid(context.Context, *users.MultipleUuids) domain.Response
	UpdateUser(context.Context, *users.UpdateUserRequest) domain.Response
	SearchContact(ctx context.Context, request *users.SearchContactRequest) domain.Response
}

func NewUsersService(users_repo repository.UsersRepositoryInterface) UsersServiceInterface {
	return &usersService{users_repo: users_repo}
}

func (us usersService) CreateUser(ctx context.Context, request *users.RegisterUser) domain.Response {
	return Response.CreateResponse(nil, us.users_repo.CreateUser(ctx, request))
}

func (us usersService) GetUserProfileByUuid(ctx context.Context, request *users.MultipleUuids) domain.Response {
	return Response.CreateResponse(us.users_repo.GetUserProfileByUuid(ctx, request))

}
func (us usersService) UpdateUser(ctx context.Context, request *users.UpdateUserRequest) domain.Response {
	return Response.CreateResponse(us.users_repo.UpdateUser(ctx, request))

}

func (us usersService) SearchContact(ctx context.Context, request *users.SearchContactRequest) domain.Response {
	return Response.CreateResponse(us.users_repo.SearchContact(ctx, request))
}
