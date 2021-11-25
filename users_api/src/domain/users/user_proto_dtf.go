package users

import pb "github.com/flydevs/chat-app-api/users-api/src/clients/rpc/users"

func (us User) Poblate_StructtoProto(pu *pb.User) {
	pu.Id = us.Id
	pu.Uuid = &pb.Uuid{Uuid: us.Uuid}
	pu.LoginUser = us.LoginUser
	pu.LoginPassword = us.LoginPassword
}

func (us *User) Poblate_PrototoStruct(pu *pb.User) {
	us.Id = pu.Id
	if pu.Uuid != nil {
		us.Uuid = pu.Uuid.Uuid
	}
	us.LoginUser = pu.LoginUser
	us.LoginPassword = pu.LoginPassword
}

func (usp *UserProfile) Poblate_PrototoStruct(pup *pb.UserProfile) {
	usp.Active = pup.Active
	usp.Phone = pup.Phone
	usp.FirstName = pup.FirstName
	usp.LastName = pup.LastName
	usp.UserName = pup.UserName
	usp.Description = pup.Description
	usp.AvatarUrl = pup.AvatarUrl
}

func (usp UserProfile) Poblate_StructtoProto(pup *pb.UserProfile) {
	pup.Active = usp.Active
	pup.Phone = usp.Phone
	pup.FirstName = usp.FirstName
	pup.LastName = usp.LastName
	pup.UserName = usp.UserName
	pup.Description = usp.Description
	pup.AvatarUrl = usp.AvatarUrl
	pup.CreatedAt = usp.CreatedAt
}

func (uap *UuidandProfile) Poblate_StructtoProto(upu *pb.UserProfileUuid) {
	uap.Profile.Poblate_StructtoProto(upu.User)
	upu.Uuid = &pb.Uuid{Uuid: uap.Uuid}
}

func (uap *UuidandProfile) Poblate_PrototoStruct(upu *pb.UserProfileUuid) {
	uap.Profile.Poblate_PrototoStruct(upu.User)
	uap.Uuid = upu.Uuid.Uuid
}

func (ru *RegisterUser) Poblate_PrototoStruct(pru *pb.RegisterUser) {
	ru.LoginInfo.Poblate_PrototoStruct(pru.LoginInfo)
	ru.ProfileInfo.Poblate_PrototoStruct(pru.ProfileInfo)
}

func (ru RegisterUser) Poblate_StructtoProto(pru *pb.RegisterUser) {
	ru.LoginInfo.Poblate_StructtoProto(pru.LoginInfo)
	ru.ProfileInfo.Poblate_StructtoProto(pru.ProfileInfo)
}

func (ups *UserProfileSlice) Poblate(direction_out bool, proto_ups []*pb.UserProfile) []*pb.UserProfile {
	if direction_out {
		var object_to_return []*pb.UserProfile
		for _, user_profile := range *ups {
			proto_user_profile := pb.UserProfile{}
			user_profile.Poblate_StructtoProto(&proto_user_profile)
			object_to_return = append(object_to_return, &proto_user_profile)
		}
		return object_to_return
	} else {
		for _, proto_user_profile := range proto_ups {
			user_profile := UserProfile{}
			user_profile.Poblate_PrototoStruct(proto_user_profile)
			*ups = append(*ups, &user_profile)
		}
		return nil
	}
}

func (ups *UserSlice) Poblate(direction_out bool, proto_ups []*pb.User) []*pb.User {
	if direction_out {
		var object_to_return []*pb.User
		for _, user := range *ups {
			proto_user := pb.User{}
			user.Poblate_StructtoProto(&proto_user)
			object_to_return = append(object_to_return, &proto_user)
		}
		return object_to_return
	} else {
		for _, proto_user := range proto_ups {
			user := User{}
			user.Poblate_PrototoStruct(proto_user)
			*ups = append(*ups, &user)
		}
		return nil
	}
}
