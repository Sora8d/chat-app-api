package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Config map[string]string

func init() {
	godotenv.Load("rest_env.env")
	Config = map[string]string{
		"PORT":              os.Getenv("PORT"),
		"MESSAGING_ADDRESS": os.Getenv("MESSAGING_ADDRESS"),
		"USERS_ADDRESS":     os.Getenv("USERS_ADDRESS"),
	}
}
