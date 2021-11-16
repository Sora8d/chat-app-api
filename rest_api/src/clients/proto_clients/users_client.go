package proto_clients

import (
	"fmt"

	"github.com/flydevs/chat-app-api/common/logger"
	proto_users "github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/users"
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"google.golang.org/grpc"
)

var users_proto_client usersProtoClient

type usersProtoClient struct {
	Client proto_users.UsersProtoInterfaceClient
	Conn   *grpc.ClientConn
}

func (upc *usersProtoClient) setConnection(c *grpc.ClientConn) {
	upc.Conn = c
}
func (upc *usersProtoClient) setClient(c proto_users.UsersProtoInterfaceClient) {
	upc.Client = c
}

func init() {
	logger.Info(fmt.Sprintf("connecting to users service with address: %s", config.Config["USERS_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	connection, err := grpc.Dial(config.Config["USERS_ADDRESS"], opts...)
	if err != nil {
		logger.Error("unable to connect to users_api", err)
		panic(err)
	}

	users_proto_client.setConnection(connection)
	users_proto_client.setClient(proto_users.NewUsersProtoInterfaceClient(connection))
}

func GetUsersClient() *usersProtoClient {
	return &users_proto_client
}
