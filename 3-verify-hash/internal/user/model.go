package user

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password" gorm:"required"`
	Name     string `json:"name"`
}

func NewUser(email, password, name string) (*User, error) {
	user := &User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
