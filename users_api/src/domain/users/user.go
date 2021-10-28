package users

import (
	pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc"
)

type User struct {
	Id   int64  `json:"id"`
	Uuid string `json:"uuid"`
}

type UserProfile struct {
	Id          int64  `json:"id,omitempty"`
	UserId      int64  `json:"user_id"`
	Active      bool   `json:"active"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	Description string `json:"description"`
}

type UuidandProfile struct {
	Uuid    string      `json:"uuid"`
	Profile UserProfile `json:"profile_info"`
}

func (us User) PoblateUser_StructtoProto(pu *pb.User) {
	pu.Id = us.Id
	pu.Uuid = &pb.Uuid{Uuid: us.Uuid}
}

func (us *User) PoblateUser_PrototoStruct(pu *pb.User) {
	us.Id = pu.Id
	us.Uuid = pu.Uuid.Uuid
}

func (usp *UserProfile) PoblateUserProfile_PrototoStruct(pup *pb.UserProfile) {
	usp.Active = pup.Active
	usp.Phone = pup.Phone
	usp.FirstName = pup.FirstName
	usp.LastName = pup.LastName
	usp.UserName = pup.UserName
	usp.Description = pup.Description
	usp.AvatarUrl = pup.AvatarUrl
}

func (usp UserProfile) PoblateUserProfile_StructtoProto(pup *pb.UserProfile) {
	pup.Active = usp.Active
	pup.Phone = usp.Phone
	pup.FirstName = usp.FirstName
	pup.LastName = usp.LastName
	pup.UserName = usp.UserName
	pup.Description = usp.Description
	pup.AvatarUrl = usp.AvatarUrl
}

func (uap *UuidandProfile) PoblateUuidProfile_StructtoProto(upu *pb.UserProfileUuid) {
	uap.Profile.PoblateUserProfile_StructtoProto(upu.User)
	upu.Uuid = &pb.Uuid{Uuid: uap.Uuid}
}

func (uap *UuidandProfile) PoblateUuidProfile_PrototoStruct(upu *pb.UserProfileUuid) {
	uap.Profile.PoblateUserProfile_PrototoStruct(upu.User)
	uap.Uuid = upu.Uuid.Uuid
}
