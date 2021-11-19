package db

import (
	"context"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/common/server_message"
	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
)

const (
	queryGetUserByUuid      = "SELECT id, uuid FROM user_table WHERE uuid=$1;"
	queryGetUserByLogin     = "SELECT id, uuid, login_user FROM user_table WHERE login_user=$1 and login_password=$2;"
	queryGetUserProfileById = "SELECT up.id, up.user_id, up.active, up.phone, up.first_name, up.last_name, up.username, up.avatar_url, up.description, to_char(up.created_at, 'YYYY-MM-DD HH24:MI:SS TZ') FROM user_profile up JOIN user_table ut on up.user_id = ut.id WHERE ut.uuid=$1;"
	queryDeleteUserByUuid   = "DELETE FROM user_table WHERE uuid=$1;"
	queryInsertUserProfile  = "INSERT INTO user_profile(user_id, phone, active, first_name, last_name, description, username, avatar_url) VALUES ($1,$2, $3, $4, $5, $6, $7, $8);"
	queryInsertUser         = "INSERT INTO user_table(login_user, login_password) VALUES($1, $2) RETURNING id, uuid, login_user;"
	queryUpdateUserProfile  = "UPDATE user_profile as up SET phone=$2, first_name=$3, last_name=$4, username=$5, avatar_url=$6, description=$7 from user_table as ut WHERE up.user_id=ut.id AND ut.uuid = $1 RETURNING up.active, up.phone, up.first_name, up.last_name, up.username, up.avatar_url, up.description, to_char(up.created_at, 'YYYY-MM-DD HH24:MI:SS TZ');"
	queryUpdateActive       = "UPDATE user_profile up SET active=$2 from user_table ut WHERE up.user_id = ut.id and ut.uuid = $1;"

//Here is where the queries are going to be
)

const (
	errUniquePhoneConstraint = "user_profile_phone_key"
	errNoRows                = "no rows in result set"
)

var (
	ctx         = context.Background()
	GoquDialect goqu.DialectWrapper
)

func init() {
	GoquDialect = goqu.Dialect("postgres")
}

type UserDbRepository interface {
	GetUserByUuid([]string) ([]*users.User, server_message.Svr_message)
	GetUserProfileById([]string) ([]*users.UserProfile, server_message.Svr_message)
	CreateUser(users.RegisterUser) server_message.Svr_message
	DeleteUser(string) server_message.Svr_message
	UpdateUserProfile(string, users.UserProfile) (*users.UserProfile, server_message.Svr_message)
	UpdateUserProfileActive(string, bool) server_message.Svr_message
	LoginUser(users.User) (*users.User, server_message.Svr_message)
}

func GetUserDbRepository() UserDbRepository {
	return &userDbRepository{}
}

type userDbRepository struct {
}

func (dbr *userDbRepository) GetUserByUuid(uuids []string) ([]*users.User, server_message.Svr_message) {
	client := postgresql.GetSession()
	query := GoquDialect.From("user_table").Select("id", "uuid").Where(goqu.Ex{"uuid": uuids})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating goqu sql in GetUserByUuid", err)
		return nil, server_message.NewInternalError()
	}
	rows, err := client.Query(toSQL)
	if err != nil {
		logger.Error("error executing sql in GetUserByUuid", err)
		return nil, server_message.NewInternalError()
	}
	user_array := []*users.User{}
	for rows.Next() {
		user := users.User{}
		if err := rows.Scan(&user.Id, &user.Uuid); err != nil {
			getErr := server_message.NewBadRequestError("error ocurred fetching the id")
			return nil, getErr
		}
		user_array = append(user_array, &user)
	}
	if len(user_array) == 0 {
		return nil, server_message.NewNotFoundError("no users were found")
	}
	return user_array, nil
}

func (dbr *userDbRepository) GetUserProfileById(uuids []string) ([]*users.UserProfile, server_message.Svr_message) {
	client := postgresql.GetSession()
	var profiles []*users.UserProfile
	query := GoquDialect.From(
		"user_profile").Select(
		"user_profile.id", "user_profile.user_id", "user_profile.active", "user_profile.phone", "user_profile.first_name", "user_profile.last_name", "user_profile.username", "user_profile.avatar_url", "user_profile.description", goqu.L("to_char(user_profile.created_at, 'YYYY-MM-DD HH24:MI:SS TZ')")).Join(
		goqu.T("user_table"), goqu.On(goqu.Ex{"user_profile.user_id": goqu.I("user_table.id")})).Where(goqu.Ex{"user_table.uuid": uuids})
	toSQL, _, err := query.ToSQL()
	if err != nil {
		logger.Error("error generating goqu sql in GetUserProfileById", err)
		return nil, server_message.NewInternalError()
	}
	rows, err := client.Query(toSQL)
	if err != nil {
		logger.Error("error executing sql in GetUserProfile", err)
		getErr := server_message.NewInternalError()
		logger.Error(getErr.GetFormatted(), err)
		return nil, getErr
	}
	for rows.Next() {
		var profile users.UserProfile
		if err := rows.Scan(&profile.Id, &profile.UserId, &profile.Active, &profile.Phone, &profile.FirstName, &profile.LastName, &profile.UserName, &profile.AvatarUrl, &profile.Description, &profile.CreatedAt); err != nil {
			getErr := server_message.NewInternalError()
			logger.Error(getErr.GetFormatted(), err)
			return nil, getErr
		}
		profiles = append(profiles, &profile)
	}
	if len(profiles) == 0 {
		return nil, server_message.NewNotFoundError("no users were found")
	}
	return profiles, nil
}

func (dbr *userDbRepository) CreateUser(uc users.RegisterUser) server_message.Svr_message {
	client := postgresql.GetSession()
	tx, err := client.Transaction()
	if err != nil {
		transErr := server_message.NewInternalError()
		logger.Error("error trying to begin a transaction in CreateUser fuction in the db_repository", err)
		return transErr
	}
	defer tx.Rollback(ctx)

	var newUser users.User
	row_with_user := tx.QueryRow(ctx, queryInsertUser, uc.LoginInfo.LoginUser, uc.LoginInfo.LoginPassword)
	if err := row_with_user.Scan(&newUser.Id, &newUser.Uuid, &newUser.LoginUser); err != nil {
		transErr := server_message.NewInternalError()
		logger.Error("error trying to create user in CreateUser fuction in the db_repository", err)
		return transErr
	}

	_, err = tx.Exec(ctx, queryInsertUserProfile, newUser.Id, uc.ProfileInfo.Phone, uc.ProfileInfo.Active, uc.ProfileInfo.FirstName, uc.ProfileInfo.LastName, uc.ProfileInfo.Description, uc.ProfileInfo.UserName, uc.ProfileInfo.AvatarUrl)
	if err != nil {
		if strings.Contains(err.Error(), errUniquePhoneConstraint) {
			transErr := server_message.NewBadRequestError("phone number already registered")
			return transErr
		}
		//Later make an if to unique costraint breaks
		transErr := server_message.NewInternalError()
		//		logger.Error("error trying to create user_profile in CreateUser fuction in the db_repository", err)
		return transErr
	}

	tx.Commit(ctx)
	return nil
}

func (dbr *userDbRepository) LoginUser(log_info users.User) (*users.User, server_message.Svr_message) {
	client := postgresql.GetSession()

	var resp_user users.User

	row := client.QueryRow(queryGetUserByLogin, log_info.LoginUser, log_info.LoginPassword)
	if err := row.Scan(&resp_user.Id, &resp_user.Uuid, &resp_user.LoginUser); err != nil {
		if err.Error() == errNoRows {
			return nil, server_message.NewBadRequestError("invalid credentials")
		}
		logger.Error("there was an error in LoginUser", err)
		return nil, server_message.NewInternalError()
	}
	return &resp_user, nil
}

func (dbr *userDbRepository) DeleteUser(uuid string) server_message.Svr_message {
	client := postgresql.GetSession()

	if err := client.Execute(queryDeleteUserByUuid, uuid); err != nil {
		//Later make an unique constraint for not found
		delErr := server_message.NewBadRequestError("there was an error deleting user with given uuid")
		return delErr
	}
	return nil
}

func (dbr *userDbRepository) UpdateUserProfile(uuid string, up users.UserProfile) (*users.UserProfile, server_message.Svr_message) {
	client := postgresql.GetSession()

	profile := users.UserProfile{}
	row := client.QueryRow(queryUpdateUserProfile, &uuid, &up.Phone, &up.FirstName, &up.LastName, &up.UserName, &up.AvatarUrl, &up.Description)
	if err := row.Scan(&profile.Active, &profile.Phone, &profile.FirstName, &profile.LastName, &profile.UserName, &profile.AvatarUrl, &profile.Description, &profile.CreatedAt); err != nil {
		upErr := server_message.NewBadRequestError("there was an error updating user")
		return nil, upErr
	}

	return &profile, nil
}

func (dbr *userDbRepository) UpdateUserProfileActive(uuid string, active bool) server_message.Svr_message {
	client := postgresql.GetSession()

	if err := client.Execute(queryUpdateActive, &uuid, &active); err != nil {
		actErr := server_message.NewInternalError()
		logger.Error(actErr.GetFormatted(), err)
		return actErr
	}
	return nil
}
