package entity

import "github.com/dgrijalva/jwt-go"

type Entity struct {
	jwt.StandardClaims
	Uuid        string `json:"uuid"`
	Permissions int    `json:"permissions"`
}
