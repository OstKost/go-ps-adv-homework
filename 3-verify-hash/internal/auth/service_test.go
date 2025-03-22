package auth_test

import (
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/user"
	"testing"
)

type MockUserRepository struct {
}

func (repository *MockUserRepository) CreateUser(u *user.User) (*user.User, error) {
	return &user.User{Email: "test@test.test"}, nil
}

func (repository *MockUserRepository) UpdateUser(u *user.User) (*user.User, error) {
	return nil, nil
}

func (repository *MockUserRepository) DeleteUser(id uint) error {
	return nil
}

func (repository *MockUserRepository) GetUserById(id uint) (*user.User, error) {
	return nil, nil
}

func (repository *MockUserRepository) GetUserByEmail(email string) (*user.User, error) {
	return nil, nil
}

func (repository *MockUserRepository) FindUsers(name, email string) (*[]user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	testUser := struct {
		Email    string
		Password string
		Name     string
	}{
		Email:    "test@test.test",
		Password: "password",
		Name:     "Test User",
	}
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(testUser.Email, testUser.Password, testUser.Name)
	if err != nil {
		t.Fatal(err)
	}
	if email != testUser.Email {
		t.Fatal("Expected email", testUser.Email, "but got", email)
	}
}
