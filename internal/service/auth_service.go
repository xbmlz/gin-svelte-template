package service

import (
	"context"
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
	opts  *options
	redis core.Redis
}

func NewAuthService(log core.Logger, conf core.Config, redis core.Redis) AuthService {
	opts := &options{
		secret:     conf.Auth.TokenSecretKey,
		expireTime: conf.Auth.TokenExpireTime,
	}
	return AuthService{
		opts,
		redis,
	}
}

func (a AuthService) GenerateToken(user *model.User) (string, error) {
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

	cmd := a.redis.Client.Set(context.Background(), "auth:"+user.Username, token, time.Duration(a.opts.expireTime)*time.Second)

	if cmd.Err() != nil {
		return "", cmd.Err()
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

	if err != nil {
		return nil, err
	}
	return token.Claims.(*dto.AuthClaims), nil
}

func (a AuthService) CleanToken(username string) error {
	cmd := a.redis.Client.Del(context.Background(), "auth:"+username)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
