package services

import (
	"github.com/flydevs/chat-app-api/common/api_errors"
	"github.com/flydevs/chat-app-api/users-api/src/domain/users"
	"github.com/flydevs/chat-app-api/users-api/src/repository/db"
)

type UsersServiceInterface interface {
	CreateUser(users.UserProfile) (*users.UuidandProfile, api_errors.Api_error)
	GetUser(string) (*users.User, api_errors.Api_error)
	GetUserProfile(string) (*users.UserProfile, api_errors.Api_error)
	DeleteUser(string) api_errors.Api_error
	UpdateUserProfile(users.UuidandProfile, bool) (*users.UserProfile, api_errors.Api_error)
	UpdateUserProfileActive(string, bool) api_errors.Api_error
}

type userService struct {
	dbRepo db.UserDbRepository
}

func NewUsersService(dbrepo db.UserDbRepository) UsersServiceInterface {
	return userService{dbRepo: dbrepo}
}

func (us userService) CreateUser(up users.UserProfile) (*users.UuidandProfile, api_errors.Api_error) {
	return us.dbRepo.CreateUser(up)
}

func (us userService) GetUser(uuid string) (*users.User, api_errors.Api_error) {
	return us.dbRepo.GetUserByUuid(uuid)
}

func (us userService) GetUserProfile(uuid string) (*users.UserProfile, api_errors.Api_error) {
	user, aErr := us.dbRepo.GetUserByUuid(uuid)
	if aErr != nil {
		return nil, aErr
	}
	result_profile := users.UserProfile{Id: user.UserProfileId}
	return us.dbRepo.GetUserProfileById(result_profile)
}

func (us userService) UpdateUserProfile(u users.UuidandProfile, partial bool) (*users.UserProfile, api_errors.Api_error) {
	var (
		uuid    = u.Uuid
		updates = u.Profile
	)
	user, aErr := us.GetUser(uuid)
	if aErr != nil {
		return nil, aErr
	}
	if partial {
		profile_with_information, aErr := us.dbRepo.GetUserProfileById(users.UserProfile{Id: user.UserProfileId})
		if aErr != nil {
			return nil, aErr
		}
		if updates.Phone != "" {
			profile_with_information.Phone = updates.Phone
		}
		if updates.FirstName != "" {
			profile_with_information.FirstName = updates.FirstName
		}
		if updates.LastName != "" {
			profile_with_information.LastName = updates.LastName
		}
		if updates.UserName != "" {
			profile_with_information.UserName = updates.UserName
		}
		if updates.AvatarUrl != "" {
			profile_with_information.AvatarUrl = updates.AvatarUrl
		}
		if updates.Description != "" {
			profile_with_information.Description = updates.Description
		}
		return us.dbRepo.UpdateUserProfile(*profile_with_information)
	}
	updates.Id = user.UserProfileId
	if updates.Phone == "" {
		return nil, api_errors.NewBadRequestError("the requested no_partial, updated brings a nil phone value")
	}
	return us.dbRepo.UpdateUserProfile(updates)
}

func (us userService) UpdateUserProfileActive(uuid string, active bool) api_errors.Api_error {
	user, aErr := us.GetUser(uuid)
	if aErr != nil {
		return aErr
	}
	return us.dbRepo.UpdateUserProfileActive(users.UserProfile{Id: user.UserProfileId, Active: active})
}

func (us userService) DeleteUser(uuid string) api_errors.Api_error {
	return us.dbRepo.DeleteUser(uuid)
}
