package entity

import "github.com/dgrijalva/jwt-go"

type Entity struct {
	jwt.StandardClaims
	Uuid        string `json:"uuid"`
	Permissions int    `json:"permissions"`
}

func NewEntity(expires int64, uuid string, permissions int) Entity {
	new_ent := Entity{StandardClaims: jwt.StandardClaims{
		ExpiresAt: expires},
		Uuid:        uuid,
		Permissions: permissions,
	}
	return new_ent
}
