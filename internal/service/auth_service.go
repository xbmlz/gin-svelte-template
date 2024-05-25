package service

import (
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/internal/model/dto"
)

type options struct {
	expireTime int64
	secret     string
}

type AuthService struct {
	opts *options
}

func NewAuthService(log core.Logger, conf core.Config) AuthService {
	opts := &options{
		secret:     conf.Auth.TokenSecretKey,
		expireTime: conf.Auth.TokenExpireTime,
	}
	return AuthService{
		opts: opts,
	}
}

func (a AuthService) GenerateToken(user model.User) (string, error) {
	expiresAt := time.Now().Add(time.Duration(a.opts.expireTime) * time.Second)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.AuthClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})
	// Sign and get the complete encoded token as a string using the key
	token, err := claims.SignedString([]byte(a.opts.secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a AuthService) ParseToken(tokenString string) (*dto.AuthClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.opts.secret), nil
	})

	switch {
	case token.Valid:
		err = nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		err = errors.New("invalid token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		err = errors.New("invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		err = errors.New("token expired")
	default:
		err = errors.New("invalid token")
	}

	if err != nil {
		return nil, err
	}
	return token.Claims.(*dto.AuthClaims), nil
}
