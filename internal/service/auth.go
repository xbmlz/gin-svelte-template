package service

import (
	"time"

	"github.com/xbmlz/gin-svelte-template/internal/config"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/pkg/token"
)

type options struct {
	expireTime int64
	secret     string
}

type AuthService struct {
	opts *options
}

func NewAuthService(logger logger.Logger, config config.Config) AuthService {
	logger.Info("AuthService created")
	opts := &options{
		secret:     config.JWT.Secret,
		expireTime: config.JWT.ExpireTime,
	}
	return AuthService{
		opts: opts,
	}
}

func (a AuthService) GenerateToken(user model.User) (string, error) {
	exp := time.Now().Add(time.Duration(a.opts.expireTime) * time.Second)
	token, err := token.GenerateToken(user.ID, a.opts.secret, exp)
	if err != nil {
		return "", err
	}
	return token, nil
}
