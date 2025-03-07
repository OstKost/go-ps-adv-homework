package auth

import (
	"errors"
	"go-ps-adv-homework/internal/user"
)

type AuthService struct {
	*user.UserRepository
}

func NewAuthService(repository *user.UserRepository) *AuthService {
	return &AuthService{UserRepository: repository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.GetUserByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	newUser, err := user.NewUser(email, password, name)
	if err != nil {
		return "", err
	}
	createdUser, err := service.UserRepository.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return createdUser.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.GetUserByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrInvalidCredentials)
	}
	validPassword := existedUser.CheckPassword(password)
	if !validPassword {
		return "", errors.New(ErrInvalidCredentials)
	}
	return existedUser.Email, nil
}
