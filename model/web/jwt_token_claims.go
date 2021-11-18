package web

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
}
