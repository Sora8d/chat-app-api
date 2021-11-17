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
	CreateUser(*users.RegisterUser) domain.Response
	LoginUser(*users.User) domain.Response
	GetUserProfileByUuid(*users.MultipleUuids) domain.Response
	UpdateUser(*users.UpdateUserRequest) domain.Response
}

func NewUsersService(users_repo repository.UsersRepositoryInterface) UsersServiceInterface {
	return &usersService{users_repo: users_repo}
}

func (us usersService) CreateUser(request *users.RegisterUser) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(nil, us.users_repo.CreateUser(ctx, request))
}

func (us usersService) LoginUser(request *users.User) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(us.users_repo.LoginUser(ctx, request))

}
func (us usersService) GetUserProfileByUuid(request *users.MultipleUuids) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(us.users_repo.GetUserProfileByUuid(ctx, request))

}
func (us usersService) UpdateUser(request *users.UpdateUserRequest) domain.Response {
	ctx := context.Background()
	return Response.CreateResponse(us.users_repo.UpdateUser(ctx, request))

}
