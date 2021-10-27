package db

import (
	"context"

	"github.com/flydevs/chat-app-api/common/api_errors"
	"github.com/flydevs/chat-app-api/common/logger"
	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/jackc/pgx/v4"
)

const (
	queryGetUserByUuid      = "SELECT id, uuid, user_profile_id FROM user_table WHERE uuid=$1;"
	queryGetUserProfileById = "SELECT id, active, phone, first_name, last_name, username, avatar_url, description FROM user_profile WHERE id=$1;"
	queryDeleteUserByUuid   = "DELETE user_table.*, user_profile.* FROM user_table JOIN user_profile ON user_table.user_profile_id = user_profile.id WHERE user_table.uuid=$1;"
	queryInsertUserProfile  = "INSERT INTO user_profile(phone, active, first_name, last_name, description, username, avatar_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, active, phone, first_name, last_name, username, avatar_url, description;"
	queryInsertUser         = "INSERT INTO user_table(user_profile_id) VALUES($1) RETURNING uuid;"
	queryUpdateUserProfile  = "UPDATE user_profile SET phone=$2, first_name=$3, last_name=$4, username=$5, avatar_url=$6, description =$7 WHERE id=$1 RETURNING active, phone, first_name, last_name, username, avatar_url, description;"
	queryUpdateActive       = "UPDATE user_profile SET active=$2 WHERE id=$1;"

//Here is where the queries are going to be
)

var (
	ctx = context.Background()
)

type UserDbRepository interface {
	GetUserByUuid(string) (*users.User, api_errors.Api_error)
	GetUserProfileById(users.UserProfile) (*users.UserProfile, api_errors.Api_error)
	CreateUser(users.UserProfile) (*users.UuidandProfile, api_errors.Api_error)
	DeleteUser(string) api_errors.Api_error
	UpdateUserProfile(users.UserProfile) (*users.UserProfile, api_errors.Api_error)
	UpdateUserProfileActive(users.UserProfile) api_errors.Api_error
}

func GetUserDbRepository() UserDbRepository {
	return &userDbRepository{}
}

type userDbRepository struct {
}

func (dbr *userDbRepository) GetUserByUuid(uuid string) (*users.User, api_errors.Api_error) {
	client := postgresql.GetSession()
	user := users.User{}
	row := client.QueryRow(queryGetUserByUuid, uuid)
	if err := row.Scan(&user.Id, &user.Uuid, &user.UserProfileId); err != nil {
		if err == pgx.ErrNoRows {
			return nil, api_errors.NewNotFoundError("no user with given Uuid")
		}
		getErr := api_errors.NewBadRequestError("error ocurred fetching the id")
		return nil, getErr
	}
	return &user, nil
}

func (dbr *userDbRepository) GetUserProfileById(profile users.UserProfile) (*users.UserProfile, api_errors.Api_error) {
	client := postgresql.GetSession()
	row := client.QueryRow(queryGetUserProfileById, &profile.Id)
	if err := row.Scan(&profile.Id, &profile.Active, &profile.Phone, &profile.FirstName, &profile.LastName, &profile.UserName, &profile.AvatarUrl, &profile.Description); err != nil {
		getErr := api_errors.NewInternalError("error ocurred obtaining profile", err)
		logger.Error(getErr.GetFormatted(), err)
		return nil, getErr
	}
	return &profile, nil
}

func (dbr *userDbRepository) CreateUser(up users.UserProfile) (*users.UuidandProfile, api_errors.Api_error) {
	client := postgresql.GetSession()
	tx, err := client.Transaction()
	if err != nil {
		transErr := api_errors.NewInternalError("error creating user", err)
		logger.Error("error trying to begin a transaction in CreateUser fuction in the db_repository", err)
		return nil, transErr
	}
	defer tx.Rollback(ctx)

	var profile users.UserProfile
	row_with_id := tx.QueryRow(ctx, queryInsertUserProfile, up.Phone, up.Active, up.FirstName, up.LastName, up.Description, up.UserName, up.AvatarUrl)
	if err := row_with_id.Scan(&profile.Id, &profile.Active, &profile.Phone, &profile.FirstName, &profile.LastName, &profile.UserName, &profile.AvatarUrl, &profile.Description); err != nil {
		//Later make an if to unique costraint breaks
		transErr := api_errors.NewBadRequestError("error creating user")
		//		logger.Error("error trying to create user_profile in CreateUser fuction in the db_repository", err)
		return nil, transErr
	}

	var uuid string
	row_with_uuid := tx.QueryRow(ctx, queryInsertUser, profile.Id)
	if err := row_with_uuid.Scan(&uuid); err != nil {
		transErr := api_errors.NewInternalError("error creating user", err)
		logger.Error("error trying to create user in CreateUser fuction in the db_repository", err)
		return nil, transErr
	}
	tx.Commit(ctx)
	result := users.UuidandProfile{Uuid: uuid, Profile: profile}
	return &result, nil
}

func (dbr *userDbRepository) DeleteUser(uuid string) api_errors.Api_error {
	client := postgresql.GetSession()

	if err := client.Execute(queryDeleteUserByUuid, uuid); err != nil {
		//Later make an unique constraint for not found
		delErr := api_errors.NewBadRequestError("there was an error deleting user with given uuid")
		return delErr
	}
	return nil
}

func (dbr *userDbRepository) UpdateUserProfile(up users.UserProfile) (*users.UserProfile, api_errors.Api_error) {
	client := postgresql.GetSession()

	profile := users.UserProfile{Id: up.Id}
	row := client.QueryRow(queryUpdateUserProfile, &up.Id, &up.Phone, &up.FirstName, &up.LastName, &up.UserName, &up.AvatarUrl, &up.Description)
	if err := row.Scan(&profile.Active, &profile.Phone, &profile.FirstName, &profile.LastName, &profile.UserName, &profile.AvatarUrl, &profile.Description); err != nil {
		upErr := api_errors.NewBadRequestError("there was an error updating user")
		return nil, upErr
	}

	return &profile, nil
}

func (dbr *userDbRepository) UpdateUserProfileActive(up users.UserProfile) api_errors.Api_error {
	client := postgresql.GetSession()

	if err := client.Execute(queryUpdateActive, &up.Id, &up.Active); err != nil {
		actErr := api_errors.NewInternalError("there was an error updating the 'active' state of this user", err)
		logger.Error(actErr.GetFormatted(), err)
		return actErr
	}
	return nil
}
