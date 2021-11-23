package config

import (
	"os"
)

var Config map[string]string

func init() {
	Config = map[string]string{
		"PORT":              os.Getenv("PORT"),
		"MESSAGING_ADDRESS": os.Getenv("MESSAGING_ADDRESS"),
		"USERS_ADDRESS":     os.Getenv("USERS_ADDRESS"),
		"OAUTH_ADDRESS":     os.Getenv("OAUTH_ADDRESS"),
	}
}
