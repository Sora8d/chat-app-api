package config

import "os"

var (
	Config map[string]string
)

func init() {
	Config = map[string]string{
		"OAUTH_KEY":     os.Getenv("OAUTH_ENV"),
		"USERS_KEY":     os.Getenv("USERS_ENV"),
		"MESSAGING_KEY": os.Getenv("MESSAGING_ENV"),
	}
}
