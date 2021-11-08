package db_ctrl

import (
	"context"

	"github.com/flydevs/chat-app-api/messaging-api/src/config"
	pgx "github.com/jackc/pgx/v4"
)

//This is a monster, im just playing around
type custom_db pgx.Conn

var (
	DBClient custom_db
)

func init() {
	datasource := config.Config["DATABASE"]
	var err error
	client, err := pgx.Connect(context.Background(), datasource)
	DBClient = custom_db(*client)
	if err != nil {
		panic(err)
	}
}

func (cd custom_db) Flush() {
	new := pgx.Conn(cd)
	_, err := new.Exec(context.Background(), "DELETE FROM message_table;")
	if err != nil {
		panic(err)
	}
	_, err = new.Exec(context.Background(), "DELETE FROM user_conversation;")
	if err != nil {
		panic(err)
	}
	_, err = new.Exec(context.Background(), "DELETE FROM conversation;")
	if err != nil {
		panic(err)
	}
}
