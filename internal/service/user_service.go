package service

import (
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/errors"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"github.com/xbmlz/gin-svelte-template/internal/repo"
	"golang.org/x/crypto/bcrypt"
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
func (s UserService) Create(user *model.User) error {
	if u, err := s.userRepo.GetByUsername(user.Username); err == nil && u != nil {
		return errors.ErrUserExists
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPwd)

	s.userRepo.Create(user)
	return nil
}

// Login user
func (s UserService) Login(username, password string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	return user, nil
}
