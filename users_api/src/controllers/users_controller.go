package controllers

import (
	"context"

	"github.com/flydevs/chat-app-api/common/api_errors"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/services"
)

type userServer struct {
	pb.UnimplementedUsersProtoInterfaceServer

	svc services.UsersServiceInterface
}

func (us userServer) GetUserByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.UserErrorResponse, error) {
	user, aErr := us.svc.GetUser(uuid.Uuid)

	if aErr != nil {
		var error_to_return pb.Error
		poblateError(aErr, &error_to_return)
		return &pb.UserErrorResponse{Error: &error_to_return}, nil
	}

	var user_to_return pb.User
	user.PoblateUser_StructtoProto(&user_to_return)

	return &pb.UserErrorResponse{User: &user_to_return}, nil
}

func (us userServer) GetUserProfileByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.UserProfileErrorResponse, error) {
	result, aErr := us.svc.GetUserProfile(uuid.Uuid)
	if aErr != nil {
		var error_to_return pb.Error
		poblateError(aErr, &error_to_return)
		return &pb.UserProfileErrorResponse{Error: &error_to_return}, nil
	}

	var user_p_to_return pb.UserProfile
	result.PoblateUserProfile_StructtoProto(&user_p_to_return)

	return &pb.UserProfileErrorResponse{User: &user_p_to_return}, nil
}

func (us userServer) CreateUser(ctx context.Context, up *pb.UserProfile) (*pb.UserProfileUuidResponse, error) {
	var user_profile users.UserProfile
	user_profile.PoblateUserProfile_PrototoStruct(up)
	result, aErr := us.svc.CreateUser(user_profile)
	if aErr != nil {
		var error_to_return pb.Error
		poblateError(aErr, &error_to_return)
		return &pb.UserProfileUuidResponse{Error: &error_to_return}, nil
	}
	//Esto esta hecho asi porque sino los errores de nil pointers son un laburo
	var uuid pb.Uuid = pb.Uuid{Uuid: result.Uuid}
	var responseContent pb.UserProfileUuid = pb.UserProfileUuid{Uuid: &uuid}
	var response pb.UserProfileUuidResponse = pb.UserProfileUuidResponse{Content: &responseContent}
	return &response, nil
}

func (us userServer) ModifyUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.Error, error) {
	var request users.UuidandProfile
	request.PoblateUuidProfile_PrototoStruct(mdur.Content)

	_, aErr := us.svc.UpdateUserProfile(request, mdur.Partial)
	if aErr != nil {
		var error_to_return pb.Error
		poblateError(aErr, &error_to_return)
		return &error_to_return, nil
	}
	return nil, nil
}

func (us userServer) DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.Error, error) {
	aErr := us.svc.DeleteUser(uuid.Uuid)
	if aErr != nil {
		var error_to_return pb.Error
		poblateError(aErr, &error_to_return)
		return &error_to_return, nil
	}
	return nil, nil
}

func GetNewUserServer(svc services.UsersServiceInterface) userServer {
	return userServer{svc: svc}
}

func poblateError(aErr api_errors.Api_error, pErr *pb.Error) {
	pErr.Status = int32(aErr.GetStatus())
	pErr.Message = aErr.GetMessage()
}
