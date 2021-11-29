package proto_clients

import (
	"fmt"

	"github.com/Sora8d/common/logger"
	proto_users "github.com/flydevs/chat-app-api/oauth-api/src/clients/rpc/login"
	"github.com/flydevs/chat-app-api/oauth-api/src/config"
	"google.golang.org/grpc"
)

var login_proto_client loginProtoClient

type loginProtoClient struct {
	Client proto_users.UsersProtoInterfaceClient
	Conn   *grpc.ClientConn
}

func (lpc *loginProtoClient) setConnection(c *grpc.ClientConn) {
	lpc.Conn = c
}
func (lpc *loginProtoClient) setClient(c proto_users.UsersProtoInterfaceClient) {
	lpc.Client = c
}
func init() {
	logger.Info(fmt.Sprintf("connecting to users service with address: %s", config.Config["USERS_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	connection, err := grpc.Dial(config.Config["USERS_ADDRESS"].(string), opts...)
	if err != nil {
		logger.Error("unable to connect to users_api", err)
		panic(err)
	}

	login_proto_client.setConnection(connection)
	login_proto_client.setClient(proto_users.NewUsersProtoInterfaceClient(connection))
}

func GetLoginClient() *loginProtoClient {
	return &login_proto_client
}
