package controllers

import (
	"context"
	"fmt"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/services"
	"google.golang.org/grpc/metadata"
)

type userServer struct {
	pb.UnimplementedUsersProtoInterfaceServer

	svc services.UsersServiceInterface
}

func (us userServer) UserLogin(ctx context.Context, u *pb.User) (*pb.UserMsgResponse, error) {
	var user_log users.User
	user_log.Poblate_PrototoStruct(u)
	res, msg := us.svc.LoginUser(user_log)
	var response pb.UserMsgResponse
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	response.Msg = &msg_to_return
	if msg.GetStatus() >= 400 {

		return &response, nil
	}
	var user_to_return pb.User
	res.Poblate_StructtoProto(&user_to_return)
	response.Users = []*pb.User{&user_to_return}
	return &response, nil
}

func (us userServer) GetUserByUuid(ctx context.Context, uuid *pb.MultipleUuids) (*pb.UserMsgResponse, error) {
	//This is just part of the oauth mock
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && md.Get("user_uuid") != nil && md.Get("admin") != nil {
		logger.Info(fmt.Sprintf("user: %s, permissions: %v", md.Get("user_uuid")[0], md.Get("admin")[0]))
	}
	var uuids []string
	for _, proto_uuids := range uuid.Uuids {
		uuids = append(uuids, proto_uuids.Uuid)
	}
	users, msg := us.svc.GetUser(uuids)

	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)

	if users != nil {
		user_to_return := users.Poblate(true, nil)
		response := pb.UserMsgResponse{Users: user_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}

}

func (us userServer) GetUserProfileByUuid(ctx context.Context, uuid *pb.MultipleUuids) (*pb.UserProfileMsgResponse, error) {
	var uuids []string
	for _, proto_uuids := range uuid.Uuids {
		uuids = append(uuids, proto_uuids.Uuid)
	}
	result, msg := us.svc.GetUserProfile(uuids)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	if result != nil {
		user_p_to_return := result.Poblate(true, nil)
		response := pb.UserProfileMsgResponse{User: user_p_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserProfileMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}
}

func (us userServer) CreateUser(ctx context.Context, ru *pb.RegisterUser) (*pb.SvrMsg, error) {
	var user_profile users.RegisterUser
	user_profile.Poblate_PrototoStruct(ru)
	msg := us.svc.CreateUser(user_profile)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	return &msg_to_return, nil
}

func (us userServer) UpdateUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.UserProfileMsgResponse, error) {
	var request users.UuidandProfile
	request.Poblate_PrototoStruct(mdur.Content)

	resp_profile, msg := us.svc.UpdateUserProfile(request, mdur.Partial)
	var msg_to_return pb.SvrMsg
	poblateMessage(msg, &msg_to_return)
	if resp_profile != nil {
		var proto_user_profile pb.UserProfile
		var slice_to_return []*pb.UserProfile
		resp_profile.Poblate_StructtoProto(&proto_user_profile)
		slice_to_return = append(slice_to_return, &proto_user_profile)

		response := pb.UserProfileMsgResponse{User: slice_to_return, Msg: &msg_to_return}
		return &response, nil
	} else {
		response := pb.UserProfileMsgResponse{Msg: &msg_to_return}
		return &response, nil
	}
}

func (us userServer) UpdateActive(ctx context.Context, req *pb.UpdateActiveRequest) (*pb.SvrMsg, error) {
	var msg_to_return pb.SvrMsg
	result_msg := us.svc.UpdateUserProfileActive(req.Uuid.Uuid, req.Active)
	poblateMessage(result_msg, &msg_to_return)
	return &msg_to_return, nil
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
