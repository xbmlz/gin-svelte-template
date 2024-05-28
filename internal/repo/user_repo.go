package repo

import (
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/model"
)

// UserRepo data access layer for user
type UserRepo struct {
	db  core.Database
	log core.Logger
}

// NewUserRepo creates a new user repository
func NewUserRepo(db core.Database, log core.Logger) UserRepo {
	return UserRepo{db: db, log: log}
}

// GetByUsername returns user by username
func (r UserRepo) GetByUsername(username string) (user *model.User, err error) {
	if err = r.db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create creates a new user
func (r UserRepo) Create(user *model.User) (err error) {
	err = r.db.DB.Create(user).Error
	return
}
