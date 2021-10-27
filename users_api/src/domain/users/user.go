package users

type User struct {
	Id            int64  `json:"id"`
	Uuid          string `json:"uuid"`
	UserProfileId int64  `json:"user_profile_id"`
}

type UserProfile struct {
	Id          int64  `json:"id,omitempty"`
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
