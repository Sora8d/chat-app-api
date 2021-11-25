package config

import (
	"os"
)

var Config map[string]string

func init() {
	Config = map[string]string{
		"PORT":              os.Getenv("PORT"),
		"DATABASE":          os.Getenv("DATABASE_URL"),
		"USERS_ADDRESS":     os.Getenv("USERS_ADDRESS"),
		"OAUTH_ADDRESS":     os.Getenv("OAUTH_ADDRESS"),
		"SECRET_KEY":        os.Getenv("SECRET_KEY"),
		"TWILIO_ACC_SID":    os.Getenv("TWILIO_ACC_SID"),
		"TWILIO_AUTH_TOKEN": os.Getenv("TWILIO_AUTH_TOKEN"),
	}
}
