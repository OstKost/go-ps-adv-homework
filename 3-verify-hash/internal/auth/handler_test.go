package auth_test

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/user"
	"go-ps-adv-homework/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func bootstrapMockDB() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, err
	}
	userRepository := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepository),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrapMockDB()
	if err != nil {
		t.Fatalf("error bootstrapping mock db: %v", err)
		return
	}
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test@test.test", "$2a$10$PSxRcr1k2DMrUqrNHCPq2OS6dknAZm32fAaDV66hnPX8BEKxcEFKa")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.test",
		Password: "test123",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)

	handler.Login()(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected %d status code, but got %d", http.StatusOK, w.Result().StatusCode)
	}
}

func TestRegisterHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrapMockDB()
	if err != nil {
		t.Fatalf("error bootstrapping mock db: %v", err)
		return
	}
	// Empty user
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "test@test.test",
		Password: "test123",
		Name:     "Test User",
	})

	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)

	handler.Register()(w, req)
	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("Expected %d status code, but got %d", http.StatusOK, w.Result().StatusCode)
	}
}
