package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"column:id;size:36;index;not null;" json:"id"`
	Username string `gorm:"column:username;size:64;not null;index;" json:"username" validate:"required"`
	Password string `gorm:"column:password;not null;" json:"password" json:"phone"`
	Email    string `gorm:"column:email;default:'';" json:"email"`
	Phone    string `gorm:"column:phone;default:'';" json:"phone"`
	Status   int    `gorm:"column:status;not null;default:0;" json:"status" validate:"required,max=1,min=-1"`
}
