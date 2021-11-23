package server

import (
	"context"
	"net/http"

	"github.com/flydevs/chat-app-api/common/server_message"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/users-api/src/controllers"
)

type userServer struct {
	pb.UnimplementedUsersProtoInterfaceServer

	ctrl controllers.UserControllerInterface
}

func (us userServer) UserLogin(ctx context.Context, u *pb.User) (*pb.UserMsgResponse, error) {
	result, err := us.ctrl.UserLogin(ctx, u)
	var response pb.UserMsgResponse
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user logged"), &msg_to_return)
	}
	response.Msg = &msg_to_return
	response.Users = []*pb.User{result}
	return &response, nil

}
func (us userServer) GetUserByUuid(ctx context.Context, uuids *pb.MultipleUuids) (*pb.UserMsgResponse, error) {
	result, err := us.ctrl.GetUserByUuid(ctx, uuids)
	var response pb.UserMsgResponse
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user retrieved"), &msg_to_return)
	}
	response.Msg = &msg_to_return
	response.Users = result
	return &response, nil
}
func (us userServer) GetUserProfileByUuid(ctx context.Context, uuids *pb.MultipleUuids) (*pb.UserProfileMsgResponse, error) {
	result, err := us.ctrl.GetUserProfileByUuid(ctx, uuids)
	var response pb.UserProfileMsgResponse
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user retrieved"), &msg_to_return)
	}
	response.Msg = &msg_to_return
	response.User = result
	return &response, nil
}
func (us userServer) CreateUser(ctx context.Context, ru *pb.RegisterUser) (*pb.SvrMsg, error) {
	err := us.ctrl.CreateUser(ctx, ru)
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user created"), &msg_to_return)
	}
	return &msg_to_return, nil
}

func (us userServer) UpdateUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.UserProfileMsgResponse, error) {
	result, err := us.ctrl.UpdateUser(ctx, mdur)
	var response pb.UserProfileMsgResponse
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user updated"), &msg_to_return)
	}
	response.Msg = &msg_to_return
	response.User = []*pb.UserProfile{result}
	return &response, nil
}

func (us userServer) UpdateActive(ctx context.Context, req *pb.UpdateActiveRequest) (*pb.SvrMsg, error) {
	err := us.ctrl.UpdateActive(ctx, req)
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "active status updated"), &msg_to_return)
	}
	return &msg_to_return, nil
}

func (us userServer) DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) (*pb.SvrMsg, error) {
	err := us.ctrl.DeleteUserByUuid(ctx, uuid)
	var msg_to_return pb.SvrMsg
	if err != nil {
		poblateMessage(err, &msg_to_return)
	} else {
		poblateMessage(server_message.NewCustomMessage(http.StatusOK, "user deleted"), &msg_to_return)
	}
	return &msg_to_return, nil
}

func GetNewUserServer(users_controller controllers.UserControllerInterface) userServer {
	return userServer{ctrl: users_controller}
}

func poblateMessage(msg server_message.Svr_message, pErr *pb.SvrMsg) {
	pErr.Status = int32(msg.GetStatus())
	pErr.Message = msg.GetMessage()
}
