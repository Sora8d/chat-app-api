package postgresql

import (
	"context"

	"github.com/flydevs/chat-app-api/messaging-api/src/config"

	"github.com/flydevs/chat-app-api/common/logger"

	pgx "github.com/jackc/pgx/v4"
)

type DbClient interface {
	QueryRow(string, ...interface{}) pgx.Row
	Query(string, ...interface{}) (pgx.Rows, error)
	Execute(string, ...interface{}) error
	Insert(string, ...interface{}) pgx.Row
	Transaction() (pgx.Tx, error)
}

type dbclient struct {
	conn *pgx.Conn
}

var (
	current_client dbclient
)

func init() {
	datasource := config.Config["DATABASE"]
	var err error
	current_client.conn, err = pgx.Connect(context.Background(), datasource)
	if err != nil {
		logger.Error("Error connecting to the database, shutting down the app", err)
		panic(err)
	}
}
func GetSession() DbClient {
	return current_client
}

func (dbcl dbclient) QueryRow(query string, args ...interface{}) pgx.Row {
	row := dbcl.conn.QueryRow(context.Background(), query, args...)
	return row
}

func (dbcl dbclient) Query(query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := dbcl.conn.Query(context.Background(), query, args...)
	return rows, err
}
func (dbcl dbclient) Execute(query string, args ...interface{}) error {
	_, err := dbcl.conn.Exec(context.Background(), query, args...)
	return err
}

func (dbcl dbclient) Insert(query string, args ...interface{}) pgx.Row {
	row := dbcl.conn.QueryRow(context.Background(), query, args...)
	return row
}

func (dbcl dbclient) Transaction() (pgx.Tx, error) {
	return dbcl.conn.Begin(context.Background())
}
