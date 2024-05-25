package dto

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	ID       string
	Username string
	jwt.RegisteredClaims
}
