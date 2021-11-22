package users

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	jwt.StandardClaims
	Uuid string `json:"uuid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
