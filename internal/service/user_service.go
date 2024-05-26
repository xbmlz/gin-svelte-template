package service

import (
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/errors"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/internal/repo"
)

// UserService user service layer
type UserService struct {
	log      core.Logger
	userRepo repo.UserRepo
}

// NewUserService create new user service
func NewUserService(log core.Logger, userRepo repo.UserRepo) UserService {
	return UserService{
		log:      log,
		userRepo: userRepo,
	}
}

// Create user
func (s UserService) Create(user *model.User) (err error) {
	u, err := s.userRepo.GetByUsername(user.Username)
	if err != nil {
		return
	}
	if u != nil {
		err = errors.ErrUserExists
		return
	}

	s.userRepo.Create(user)
	return
}
