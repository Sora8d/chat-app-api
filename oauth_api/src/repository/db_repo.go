package repository

import (
	"fmt"

	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
	"github.com/doug-martin/goqu/v9"
	"github.com/flydevs/chat-app-api/oauth-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/oauth-api/src/domain/refresh_token"
)

var GoquDialect goqu.DialectWrapper

func init() {
	GoquDialect = goqu.Dialect("postgres")
}

type dbRepository struct {
}

type DbRepositoryInterface interface {
	AddNewTokenFamily(ref_token *refresh_token.RefreshToken) server_message.Svr_message
	AddNewToken(ref_token *refresh_token.RefreshToken) server_message.Svr_message
	CleanTokens(string) error
	BlockTokenFamily(family_id int64) server_message.Svr_message
	BlockFamiliesByUser(user_uuid string) server_message.Svr_message
	CheckToken(ref_token *refresh_token.RefreshToken) (*refresh_token.RefreshToken, bool, server_message.Svr_message)
}

func NewDbRepository() DbRepositoryInterface {
	return &dbRepository{}
}

func (dbrepo dbRepository) AddNewTokenFamily(ref_token *refresh_token.RefreshToken) server_message.Svr_message {
	client := postgresql.GetSession()
	query := GoquDialect.Insert("family_tokens").Cols("user", "expires").Vals([]interface{}{ref_token.Uuid, goqu.L(fmt.Sprintf("to_timestamp(%d)", ref_token.ExpiresAt))}).Returning("id")
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql AddNewTokenFamily function in db_repo", err)
		return server_message.NewInternalError()
	}
	row := client.Insert(toSQL)
	if err := row.Scan(&ref_token.Family); err != nil {
		logger.Error("error scanning AddNewTokenFamily function in db_repo", err)
		return server_message.NewInternalError()
	}
	return nil

}

func (dbrepo dbRepository) AddNewToken(ref_token *refresh_token.RefreshToken) server_message.Svr_message {
	client := postgresql.GetSession()
	query := GoquDialect.Update("family_tokens").Set(goqu.Record{"expires": goqu.L(fmt.Sprintf("to_timestamp(%d)", ref_token.ExpiresAt))}).Where(goqu.Ex{"id": ref_token.Family})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql AddNewToken function in db_repo", err)
		return server_message.NewInternalError()
	}
	err = client.Execute(toSQL)
	if err != nil {
		logger.Error("error executing query in AddNewToken function in db_repo", err)
		return server_message.NewInternalError()
	}
	return nil
}

func (dbrepo dbRepository) BlockTokenFamily(family_id int64) server_message.Svr_message {
	client := postgresql.GetSession()
	query := GoquDialect.Update("family_tokens").Set(goqu.Record{"blocked": true}).Where(goqu.Ex{"id": family_id})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql BlockTokenFamily function in db_repo", err)
		return server_message.NewInternalError()
	}
	err = client.Execute(toSQL)
	if err != nil {
		logger.Error("error executing query in BlockTokenFamily function in db_repo", err)
		return server_message.NewInternalError()
	}
	return nil
}

func (dbrepo dbRepository) BlockFamiliesByUser(user_uuid string) server_message.Svr_message {
	client := postgresql.GetSession()
	query := GoquDialect.Update("family_tokens").Set(goqu.Record{"blocked": true}).Where(goqu.Ex{"user": user_uuid})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql BlockFamiliesByUser function in db_repo", err)
		return server_message.NewInternalError()
	}
	err = client.Execute(toSQL)
	if err != nil {
		logger.Error("error executing query in BlockFamiliesByUser function in db_repo", err)
		return server_message.NewInternalError()
	}
	return nil
}

func (dbrepo dbRepository) CheckToken(ref_token *refresh_token.RefreshToken) (*refresh_token.RefreshToken, bool, server_message.Svr_message) {
	client := postgresql.GetSession()
	query := GoquDialect.From("family_tokens").Select("user", goqu.L("date_part('epoch',expires)"), "blocked").Where(goqu.Ex{"id": ref_token.Family})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql CheckToken function in db_repo", err)
		return nil, false, server_message.NewInternalError()
	}
	row := client.QueryRow(toSQL)
	var dbToken refresh_token.RefreshToken
	dbToken.Family = ref_token.Family
	var blocked bool
	if err := row.Scan(&dbToken.Uuid, &dbToken.ExpiresAt, &blocked); err != nil {
		logger.Error("error executing query in CheckToken function in db_repo", err)
		return nil, false, server_message.NewInternalError()
	}
	return &dbToken, blocked, nil
}

func (dbrepo dbRepository) CleanTokens(interval string) error {
	client := postgresql.GetSession()
	query := GoquDialect.Delete("family_tokens").Where(goqu.Ex{"expires": goqu.Op{"lt": goqu.L(fmt.Sprintf("(NOW() - interval '%s')", interval))}})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating sql CleanTokens function in db_repo", err)
		return err
	}
	err = client.Execute(toSQL)
	if err != nil {
		logger.Error("error executing query in CleanTokens function in db_repo", err)
		return err
	}
	return nil
}
