package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone     string     `json:"phone" gorm:"uniqueIndex"`
	Name      *string    `json:"name" gorm:"size:100"`
	Birthdate *time.Time `json:"birthdate" gorm:"type:date"`
}

func NewUser(phone string, name *string, birthdate *time.Time) *User {
	return &User{
		Phone:     phone,
		Name:      name,
		Birthdate: birthdate,
	}
}
