package messaging_client

import (
	"fmt"

	"github.com/flydevs/chat-app-api/common/logger"
	proto_messaging "github.com/flydevs/chat-app-api/rest-api/src/clients/rpc/messaging"
	"github.com/flydevs/chat-app-api/rest-api/src/config"
	"google.golang.org/grpc"
)

var messaging_proto_client messagingProtoClient

type messagingProtoClient struct {
	Client proto_messaging.MessagingProtoInterfaceClient
	Conn   *grpc.ClientConn
}

func (upc *messagingProtoClient) setConnection(c *grpc.ClientConn) {
	upc.Conn = c
}
func (upc *messagingProtoClient) setClient(c proto_messaging.MessagingProtoInterfaceClient) {
	upc.Client = c
}
func init() {
	logger.Info(fmt.Sprintf("connecting to messaging service with address: %s", config.Config["MESSAGING_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	connection, err := grpc.Dial(config.Config["MESSAGING_ADDRESS"], opts...)
	if err != nil {
		logger.Error("unable to connect to messaging_api", err)
		panic(err)
	}

	messaging_proto_client.setConnection(connection)
	messaging_proto_client.setClient(proto_messaging.NewMessagingProtoInterfaceClient(connection))
}

func GetMessagingClient() messagingProtoClient {
	return messaging_proto_client
}
