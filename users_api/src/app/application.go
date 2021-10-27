package app

import (
	"fmt"
	"net"

	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/users-api/src/config"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
	"github.com/flydevs/chat-app-api/users-api/src/services"
)

var usersService services.UsersServiceInterface

func StartApp() {
	postgresql.DbInit()
	usersService = services.NewUsersService(db.GetUserDbRepository())
	conn, err := net.Listen("tcp", config.Config["PORT"])
	fmt.Sprintln(conn)
	if err != nil {
		panic(err)
	}
	/*
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		fmt.Sprint(pb.User{})
	*/
}
