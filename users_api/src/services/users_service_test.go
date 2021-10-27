package services

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/flydevs/chat-app-api/users-api/src/clients/postgresql"
	"github.com/flydevs/chat-app-api/users-api/src/config"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
)

const (
	querydroptabletestuserprofile   = "DROP TABLE IF EXISTS user_profile;"
	querydroptabletestuser          = "DROP TABLE IF EXISTS user_table;"
	querycreatetabletestuserprofile = `CREATE TABLE user_profile(
		id BIGSERIAL PRIMARY KEY,
		active bool not null default true,
		phone varchar unique not null,
		first_name varchar,
		last_name varchar,
		username varchar not null,
		avatar_url varchar,
		description text
	);`
	querycreatetabletestuser = `CREATE TABLE user_table(
		id BIGSERIAL PRIMARY KEY,
		uuid uuid DEFAULT uuid_generate_v4 (),
		user_profile_id bigint unique references user_profile(id)
	);
	`
)

var (
	create_user_profile = users.UserProfile{
		Active:   true,
		Phone:    "376 4291930",
		UserName: "376 4291930",
	}

	uuid string

	client postgresql.DbClient

	svc UsersServiceInterface
)

func init() {
	config.Config["DATABASE"] = "postgres://test:123@localhost:5433/postgres"
	postgresql.DbInit()
	client = postgresql.GetSession()
	svc = NewUsersService(db.GetUserDbRepository())
}

func TestMain(m *testing.M) {

	if err := client.Execute(querydroptabletestuser); err != nil {
		panic(err)
	}
	if err := client.Execute(querydroptabletestuserprofile); err != nil {
		panic(err)
	}
	if err := client.Execute(querycreatetabletestuserprofile); err != nil {
		panic(err)
	}
	if err := client.Execute(querycreatetabletestuser); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestCreateUser(t *testing.T) {
	result, aErr := svc.CreateUser(create_user_profile)
	if aErr != nil {
		t.Error(aErr.GetFormatted(), aErr.GetError())
		t.Log(aErr.GetFormatted(), aErr.GetError())
		return
	}

	uuid = result.Uuid

	tested_profile := result.Profile
	create_user_profile.Id = tested_profile.Id
	if tested_profile != create_user_profile {
		fmt.Printf("%+v", create_user_profile)
		fmt.Printf("%+v", tested_profile)
		t.Error("tested_profile and original profile present differences")
		return
	}

	_, aErr2 := svc.CreateUser(create_user_profile)
	if aErr2 == nil {
		t.Error(aErr)
		return
	}
}

var (
	getUser_user_profile = users.UserProfile{
		Active:   true,
		Phone:    "376 4291934",
		UserName: "376 4291934",
	}
)

func TestGetUser(t *testing.T) {
	user, _ := svc.CreateUser(getUser_user_profile)

	result, aErr := svc.GetUser(user.Uuid)
	if aErr != nil {
		t.Error(aErr)
	}

	if result.UserProfileId != user.Profile.Id {
		t.Error("user doesnt match with his user_profile")
	}

	_, aErr2 := svc.GetUser("123fakeuuid")
	if aErr2 == nil {
		t.Error("There should be an error in getuser function with invalid uuid")
		return
	}
	if aErr2.GetStatus() != http.StatusBadRequest {
		t.Error("Wrong error")
		return
	}

	_, aErr3 := svc.GetUser("123e4567-e89b-12d3-a456-426614174000")
	if aErr3 == nil {
		t.Error("There should be an error in getuser function with invalid uuid")
		return
	}
	if aErr3.GetStatus() != http.StatusNotFound {
		t.Error("Wrong error")
		return
	}
}

var (
	getUserProfile_user_profile = users.UserProfile{
		Active:   true,
		Phone:    "376 4291931",
		UserName: "376 4291931",
	}
)

func TestGetUserProfile(t *testing.T) {
	new_uuid_user, _ := svc.CreateUser(getUserProfile_user_profile)

	result, aErr := svc.GetUserProfile(new_uuid_user.Uuid)
	if aErr != nil {
		t.Error(aErr)
		return
	}

	if *result != new_uuid_user.Profile {
		fmt.Printf("%+v", *result)
		fmt.Printf("%+v", new_uuid_user.Profile)
		t.Error("results should have been equal")
	}
}

var (
	updateUserProfile_original_no_partial = users.UserProfile{
		Active:    true,
		Phone:     "376 4291932",
		UserName:  "376 4291932",
		FirstName: "Jhon",
	}

	updateUserProfile_original_partial = users.UserProfile{
		Active:    true,
		Phone:     "376 4291933",
		UserName:  "376 4291933",
		FirstName: "Jhon",
	}

	updateUserProfile_changes_no_partial = users.UserProfile{
		Active:   true,
		Phone:    "376 4291932",
		UserName: "376 4291932",
		LastName: "Jhonson",
	}

	updateUserProfile_changes_no_partial_invalid = users.UserProfile{
		LastName: "Jhonson",
	}

	updateUserProfile_changes_partial = users.UserProfile{
		LastName: "Jhonson",
	}
)

func TestUpdateUser(t *testing.T) {
	no_partial_uuid_user, _ := svc.CreateUser(updateUserProfile_original_no_partial)
	changes_no_partial := users.UuidandProfile{Uuid: no_partial_uuid_user.Uuid, Profile: updateUserProfile_changes_no_partial}
	changes_no_partial_invalid := users.UuidandProfile{Uuid: no_partial_uuid_user.Uuid, Profile: updateUserProfile_changes_no_partial_invalid}

	_, aErr := svc.UpdateUserProfile(changes_no_partial_invalid, false)
	if aErr == nil {
		t.Error("there should be a unique constraint error")
		return
	}

	result, aErr := svc.UpdateUserProfile(changes_no_partial, false)
	if aErr != nil {
		t.Error(aErr)
		return
	}

	if *result == no_partial_uuid_user.Profile {
		t.Error("both user_profile objects should be different")
		return
	}
	no_partial_uuid_user.Profile.FirstName = ""
	no_partial_uuid_user.Profile.LastName = "Jhonson"

	if *result != no_partial_uuid_user.Profile {
		t.Error("both user_profile objects should be equal")
		return
	}

	partial_uuid_user, _ := svc.CreateUser(updateUserProfile_original_partial)
	changes_partial := users.UuidandProfile{Uuid: partial_uuid_user.Uuid, Profile: updateUserProfile_changes_partial}

	result2, aErr2 := svc.UpdateUserProfile(changes_partial, true)
	if aErr2 != nil {
		t.Error(aErr2)
		return
	}

	if *result2 == partial_uuid_user.Profile {
		t.Error("both user_profile objects should be different")
		return
	}
	partial_uuid_user.Profile.LastName = "Jhonson"
	if *result2 != partial_uuid_user.Profile {
		t.Error("both user_profile objects should be equal")
		return
	}

}
