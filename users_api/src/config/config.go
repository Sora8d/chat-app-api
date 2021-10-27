package config

import (
	"os"
)

var Config map[string]string

func init() {
	Config = map[string]string{
		"DATABASE": os.Getenv("DATABASE_URL"),
		"PORT":     os.Getenv("PORT"),
	}
}
