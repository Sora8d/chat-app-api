package refresh_token

import "github.com/dgrijalva/jwt-go"

type RefreshToken struct {
	jwt.StandardClaims
	Family      int64  `json:"family"`
	Uuid        string `json:"uuid"`
	Permissions int    `json:"permissions"`
}

func (token1 RefreshToken) Equal(token2 RefreshToken) bool {
	return token1.ExpiresAt == token2.ExpiresAt && token1.Family == token2.Family && token1.Uuid == token2.Uuid && token1.Permissions == token2.Permissions
}

func (token *RefreshToken) SetExpiration(expires int64) {
	token.ExpiresAt = expires
}

func NewRefreshToken(exp, family_id int64, uuid string, permissions int) RefreshToken {
	new_token := RefreshToken{Family: family_id, Uuid: uuid, Permissions: permissions}
	new_token.StandardClaims.ExpiresAt = exp
	return new_token
}
