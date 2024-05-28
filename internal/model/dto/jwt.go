package dto

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}
