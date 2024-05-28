package model

type User struct {
	BaseModel
	Username string `gorm:"not null;" json:"username" validate:"required"`
	Password string `gorm:"not null;" json:"password"`
	Email    string `gorm:"default:'';" json:"email"`
	Phone    string `gorm:"default:'';" json:"phone"`
	Status   int    `gorm:"not null;default:0;" json:"status" validate:"required,max=1,min=-1"`
}

type UserRegister struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
