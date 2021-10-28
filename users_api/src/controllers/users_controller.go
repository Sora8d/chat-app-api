package controllers

import (
	"context"

	"github.com/flydevs/chat-app-api/common/server_message"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/services"
)

type userServer struct {
	pb.UnimplementedUsersProtoInterfaceServer

	svc services.UsersServiceInterface
}

func (us userServer) GetUserByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.UserMsgResponse, error) {
	user, msg := us.svc.GetUser(uuid.Uuid)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)

	if user != nil {
		var user_to_return pb.User
		user.PoblateUser_StructtoProto(&user_to_return)
		response := pb.UserMsgResponse{User: &user_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}

}

func (us userServer) GetUserProfileByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.UserProfileMsgResponse, error) {
	result, msg := us.svc.GetUserProfile(uuid.Uuid)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	if result != nil {
		var user_p_to_return pb.UserProfile
		result.PoblateUserProfile_StructtoProto(&user_p_to_return)
		response := pb.UserProfileMsgResponse{User: &user_p_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserProfileMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}

}

func (us userServer) CreateUser(ctx context.Context, up *pb.UserProfile) (*pb.UuidResponse, error) {
	var user_profile users.UserProfile
	user_profile.PoblateUserProfile_PrototoStruct(up)
	uuid, msg := us.svc.CreateUser(user_profile)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	var uuidpbobject pb.Uuid
	if uuid != nil {
		uuidpbobject = pb.Uuid{Uuid: *uuid}
		var response pb.UuidResponse = pb.UuidResponse{Uuid: &uuidpbobject, Msg: &msg_to_return}
		return &response, nil
	} else {
		var response pb.UuidResponse = pb.UuidResponse{Msg: &msg_to_return}
		return &response, nil
	}

}

func (us userServer) UpdateUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.UserProfileMsgResponse, error) {
	var request users.UuidandProfile
	request.PoblateUuidProfile_PrototoStruct(mdur.Content)

	resp_profile, msg := us.svc.UpdateUserProfile(request, mdur.Partial)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	if resp_profile != nil {
		var user_to_return pb.UserProfile
		resp_profile.PoblateUserProfile_StructtoProto(&user_to_return)

		response := pb.UserProfileMsgResponse{User: &user_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserProfileMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}

}

func (us userServer) DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.SvrMsg, error) {
	msg := us.svc.DeleteUser(uuid.Uuid)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	return &msg_to_return, nil
}

func GetNewUserServer(svc services.UsersServiceInterface) userServer {
	return userServer{svc: svc}
}

func poblateMessage(msg server_message.Svr_message, pErr *pb.SvrMsg) {
	pErr.Status = int32(msg.GetStatus())
	pErr.Message = msg.GetMessage()
}
