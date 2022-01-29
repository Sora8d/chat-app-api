package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/users-api/src/domain/oauth"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/services"
	"google.golang.org/grpc/metadata"
)

type userController struct {
	pb.UnimplementedUsersProtoInterfaceServer

	svc services.UsersServiceInterface
}

type UserControllerInterface interface {
	UserLogin(ctx context.Context, u *pb.User) (*pb.User, server_message.Svr_message)
	GetUserByUuid(ctx context.Context, uuid *pb.MultipleUuids) ([]*pb.User, server_message.Svr_message)
	GetUserProfileByUuid(ctx context.Context, uuid *pb.MultipleUuids) ([]*pb.UserProfile, server_message.Svr_message)
	CreateUser(ctx context.Context, ru *pb.RegisterUser) server_message.Svr_message
	UpdateUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.UserProfile, server_message.Svr_message)
	UpdateActive(ctx context.Context, req *pb.UpdateActiveRequest) server_message.Svr_message
	DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) server_message.Svr_message

	SearchContact(context.Context, *pb.SearchContactRequest) ([]*pb.UserProfile, server_message.Svr_message)
}

func (us userController) UserLogin(ctx context.Context, u *pb.User) (*pb.User, server_message.Svr_message) {
	var user_log users.User
	user_log.Poblate_PrototoStruct(u)
	res, err := us.svc.LoginUser(user_log)

	var response pb.UserMsgResponse
	var msg_to_return pb.SvrMsg
	response.Msg = &msg_to_return
	if err != nil {
		return nil, err
	}
	var user_to_return pb.User
	res.Poblate_StructtoProto(&user_to_return)
	return &user_to_return, nil
}

func (us userController) GetUserByUuid(ctx context.Context, uuid *pb.MultipleUuids) ([]*pb.User, server_message.Svr_message) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && md.Get("user_uuid") != nil && md.Get("admin") != nil {
		logger.Info(fmt.Sprintf("user: %s, permissions: %v", md.Get("user_uuid")[0], md.Get("admin")[0]))
	}
	var uuids []string
	for _, proto_uuids := range uuid.Uuids {
		uuids = append(uuids, proto_uuids.Uuid)
	}
	users, err := us.svc.GetUser(uuids)
	if err != nil {
		return nil, err
	}
	user_to_return := users.Poblate(true, nil)
	return user_to_return, nil
}

func (us userController) GetUserProfileByUuid(ctx context.Context, uuid *pb.MultipleUuids) ([]*pb.UserProfile, server_message.Svr_message) {
	var uuids []string
	for _, proto_uuids := range uuid.Uuids {
		uuids = append(uuids, proto_uuids.Uuid)
	}
	result, err := us.svc.GetUserProfile(uuids)
	if err != nil {
		return nil, err
	}
	user_p_to_return := result.Poblate(true, nil)
	return user_p_to_return, nil
}

func (us userController) CreateUser(ctx context.Context, ru *pb.RegisterUser) server_message.Svr_message {
	var user_profile users.RegisterUser
	user_profile.Poblate_PrototoStruct(ru)
	err := us.svc.CreateUser(user_profile)
	return err
}

//update oauth stuff
func (us userController) UpdateUser(ctx context.Context, mdur *pb.UpdateUserRequest) (*pb.UserProfile, server_message.Svr_message) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("error reading metadata", errors.New("error in UpdateUser, no metadata"))
		return nil, server_message.NewInternalError()
	}
	at_uuid := md.Get("user_uuid")
	if at_uuid == nil {
		logger.Error("error reading metadata", errors.New("error in UpdateUser, uuid is nil"))
		return nil, server_message.NewInternalError()
	}
	if mdur.Content.Uuid.Uuid != at_uuid[0] {
		return nil, MessageBadPermission()
	}
	var request users.UuidandProfile
	request.Poblate_PrototoStruct(mdur.Content)

	resp_profile, err := us.svc.UpdateUserProfile(request, mdur.Partial)
	if err != nil {
		return nil, err
	}
	var proto_user_profile pb.UserProfile
	resp_profile.Poblate_StructtoProto(&proto_user_profile)
	return &proto_user_profile, nil
}

func (us userController) SearchContact(ctx context.Context, queries *pb.SearchContactRequest) ([]*pb.UserProfile, server_message.Svr_message) {
	md, aErr := validateContext(ctx)
	if aErr != nil {
		return nil, aErr
	}
	user_uuid, aErr := fetchUuid(md)
	if aErr != nil {
		return nil, aErr
	}

	if strings.TrimSpace(queries.Query) == "" {
		return nil, server_message.NewBadRequestError("queries cant be blank")
	}
	profiles, err := us.svc.SearchContact(queries.Query, *user_uuid, queries.ExcludeUuids)
	if err != nil {
		return nil, err
	}

	users_to_return := profiles.Poblate(true, nil)
	return users_to_return, nil
}

func (us userController) UpdateActive(ctx context.Context, req *pb.UpdateActiveRequest) server_message.Svr_message {
	result_msg := us.svc.UpdateUserProfileActive(req.Uuid.Uuid, req.Active)
	return result_msg
}

func (us userController) DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) server_message.Svr_message {
	msg := us.svc.DeleteUser(uuid.Uuid)
	return msg
}

func GetNewUserController(svc services.UsersServiceInterface) UserControllerInterface {
	return userController{svc: svc}
}

func MessageBadPermission() server_message.Svr_message {
	return server_message.NewCustomMessage(http.StatusUnauthorized, "unauhorized")
}

func validateContext(ctx context.Context) (metadata.MD, server_message.Svr_message) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("error reading metadata", errors.New("error in validatecontext, no metadata"))
		return nil, server_message.NewInternalError()
	}
	at_err := md.Get("status")
	if at_err == nil {
		logger.Error("error reading metadata", errors.New("error in validatecontext, status is nil"))
		return nil, server_message.NewInternalError()
	}
	return md, oauth.GetError(at_err[0])
}

func fetchUuid(md metadata.MD) (*string, server_message.Svr_message) {
	at_uuid := md.Get("uuid")
	if at_uuid == nil {
		logger.Error("error reading metadata", errors.New("error in UpdateUser, uuid is nil"))
		return nil, server_message.NewInternalError()
	}
	return &at_uuid[0], nil
}
