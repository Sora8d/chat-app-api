package users

import (
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	jwt.StandardClaims
	Uuid string
}
