package users_client

import (
	"context"
	"fmt"

	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
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
	GetUser(string) (*users.User, server_message.Svr_message)
}

type userProtoClient struct {
	client pb.UsersProtoInterfaceClient
	conn   *grpc.ClientConn
}

func init() {
	logger.Info(fmt.Sprintf("connecting to users_repository with address:%s", config.Config["USERS_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	c, err := grpc.Dial(config.Config["USERS_ADDRESS"], opts...)
	if err != nil {
		panic(err)
	}

	u_proto.client = pb.NewUsersProtoInterfaceClient(c)
	u_proto.conn = c
}

func (upc *userProtoClient) GetUser(uuid string) (*users.User, server_message.Svr_message) {
	pbuuid := pb.Uuid{Uuid: uuid}
	ctx := context.Background()
	user_msg_response, err := upc.client.GetUserByUuid(ctx, &pbuuid)
	if err != nil {
		logger.Error("Error response from users api", err)
		return nil, server_message.NewInternalError()
	}
	msg := poblateMsgfromProto(user_msg_response.Msg)
	if msg.GetStatus() >= 400 {
		return nil, msg
	}
	user := poblateUserfromProto(user_msg_response.User)
	return &user, nil
}

func poblateMsgfromProto(pbmsg *pb.SvrMsg) server_message.Svr_message {
	return server_message.NewCustomMessage(int(pbmsg.Status), pbmsg.Message)
}

func poblateUserfromProto(pbuser *pb.User) users.User {
	var user users.User
	user.Id = pbuser.Id
	user.Uuid = pbuser.Uuid.Uuid
	return user
}
