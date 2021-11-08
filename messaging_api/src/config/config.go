package config

import (
	"os"
)

var Config map[string]string

func init() {
	Config = map[string]string{
		"PORT":          os.Getenv("PORT"),
		"DATABASE":      os.Getenv("DATABASE_URL"),
		"USERS_ADDRESS": os.Getenv("USERS_ADDRESS"),
	}
}
