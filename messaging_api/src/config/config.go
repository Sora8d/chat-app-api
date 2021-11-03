package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Config map[string]string

func init() {
	godotenv.Load("/home/migablo/Documents/Work/flydevs_chat_api_messaging/messaging_api/test_env.env")
	Config = map[string]string{
		"DATABASE":      os.Getenv("DATABASE_URL"),
		"USERS_ADDRESS": os.Getenv("USERS_ADDRESS"),
	}
}
