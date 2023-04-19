package api

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthtokenClaims struct {
	TokenUUID string `json:"token_uuid"`
	UserID    string `json:"user_id"`
	jwt.StandardClaims
}
