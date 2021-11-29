package config

import (
	"os"
	"strconv"
)

var (
	Config map[string]interface{}
)

func init() {
	token_exp_int, err := strconv.Atoi(os.Getenv("ACCESSTOKEN_EXPIRATION"))
	if err != nil {
		panic(err)
	}
	refresh_exp_int, err := strconv.Atoi(os.Getenv("REFRESH_EXPIRATION"))
	if err != nil {
		panic(err)
	}
	Config = map[string]interface{}{
		"PORT":                   os.Getenv("PORT"),
		"DATABASE":               os.Getenv("DATABASE"),
		"SECRET_KEY":             os.Getenv("SECRET_KEY"),
		"USERS_ADDRESS":          os.Getenv("USERS_ADDRESS"),
		"OAUTH_KEY":              os.Getenv("OAUTH_KEY"),
		"USERS_KEY":              os.Getenv("USERS_KEY"),
		"MESSAGING_KEY":          os.Getenv("MESSAGING_KEY"),
		"REST_KEY":               os.Getenv("REST_KEY"),
		"ACCESSTOKEN_EXPIRATION": token_exp_int,
		"REFRESH_EXPIRATION":     refresh_exp_int,
	}
}
