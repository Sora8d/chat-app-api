package config

import "os"

var (
	Config map[string]string
)

func init() {
	Config = map[string]string{
		"PORT":          os.Getenv("PORT"),
		"SECRET_KEY":    os.Getenv("SECRET_KEY"),
		"USERS_ADDRESS": os.Getenv("USERS_ADDRESS"),
		"OAUTH_KEY":     os.Getenv("OAUTH_KEY"),
		"USERS_KEY":     os.Getenv("USERS_KEY"),
		"MESSAGING_KEY": os.Getenv("MESSAGING_KEY"),
		"REST_KEY":      os.Getenv("REST_KEY"),
	}
}
