package di

import "go-ps-adv-homework/internal/user"

type IStatRepository interface {
	AddClick(linkId uint)
}

type IUserRepository interface {
	CreateUser(user *user.User) (*user.User, error)
	UpdateUser(user *user.User) (*user.User, error)
	DeleteUser(id uint) error
	GetUserById(id uint) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	FindUsers(name, email string) (*[]user.User, error)
}
