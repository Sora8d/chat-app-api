package config

import (
	"os"
)

var Config map[string]string

func init() {
	//godotenv.Load("test_env_messaging.env")
	Config = map[string]string{
		"PORT":          os.Getenv("PORT"),
		"DATABASE":      os.Getenv("DATABASE_URL"),
		"USERS_ADDRESS": os.Getenv("USERS_ADDRESS"),
	}
}
