package users_client

import (
	"context"
	"fmt"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	pb "github.com/flydevs/chat-app-api/messaging-api/src/clients/rpc/users_client"
	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	"github.com/flydevs/chat-app-api/messaging-api/src/domain/users"
	"google.golang.org/grpc"
)

var u_proto userProtoClient

func GetUsersProtoClient() UserProtoClientRepository {
	return &u_proto
}

//Later add secret
type UserProtoClientRepository interface {
	GetUser(context.Context, []string) ([]*users.User, server_message.Svr_message)
}

type userProtoClient struct {
	client pb.UsersProtoInterfaceClient
	conn   *grpc.ClientConn
}

func (upc *userProtoClient) setConnection(c *grpc.ClientConn) {
	upc.conn = c
}
func (upc *userProtoClient) setClient(c pb.UsersProtoInterfaceClient) {
	upc.client = c
}

func init() {
	logger.Info(fmt.Sprintf("connecting to users_repository with address:%s", config.Config["USERS_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	c, err := grpc.Dial(config.Config["USERS_ADDRESS"], opts...)
	if err != nil {
		panic(err)
	}

	u_proto.setConnection(c)
	u_proto.setClient(pb.NewUsersProtoInterfaceClient(c))
}

func (upc *userProtoClient) GetUser(ctx context.Context, uuids []string) ([]*users.User, server_message.Svr_message) {
	proto_uuids := pb.MultipleUuids{}
	for _, uuid := range uuids {
		proto_uuid := pb.Uuid{Uuid: uuid}
		proto_uuids.Uuids = append(proto_uuids.Uuids, &proto_uuid)
	}
	user_msg_response, err := upc.client.GetUserByUuid(ctx, &proto_uuids)
	if err != nil {
		logger.Error("Error response from users api", err)
		return nil, server_message.NewInternalError()
	}
	msg := poblateMsgfromProto(user_msg_response.Msg)
	if msg.GetStatus() >= 400 {
		return nil, msg
	}
	user_array := poblateUserfromProto(user_msg_response.Users)
	return user_array, nil
}

func poblateMsgfromProto(pbmsg *pb.SvrMsg) server_message.Svr_message {
	return server_message.NewCustomMessage(int(pbmsg.Status), pbmsg.Message)
}

func poblateUserfromProto(proto_users []*pb.User) []*users.User {
	user_array := []*users.User{}
	for _, proto_user := range proto_users {
		var user users.User
		user.Id = proto_user.Id
		user.Uuid = proto_user.Uuid.Uuid
		user_array = append(user_array, &user)
	}
	return user_array
}
