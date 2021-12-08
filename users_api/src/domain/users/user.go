package users

type User struct {
	Id            int64  `json:"id,omitempty"`
	Uuid          string `json:"uuid"`
	LoginUser     string `json:"login_user"`
	LoginPassword string `json:"login_password,omitempty"`
}

type UserSlice []*User

type UserProfile struct {
	Id          int64  `json:"id,omitempty"`
	Uuid        string `json:"uuid"`
	UserId      int64  `json:"user_id"`
	Active      bool   `json:"active"`
	Phone       string `json:"phone"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"username"`
	AvatarUrl   string `json:"avatar_url"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type UserProfileSlice []*UserProfile

type UuidandProfile struct {
	Uuid    string      `json:"uuid"`
	Profile UserProfile `json:"profile_info"`
}
