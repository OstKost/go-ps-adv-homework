package main

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"go-ps-adv-homework/internal/auth"
	"go-ps-adv-homework/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//err = db.AutoMigrate(
	//	&link.Link{},
	//	&stat.Stat{},
	//	&user.User{},
	//)
	//if err != nil {
	//	panic(err)
	//}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "test@test.test",
		Password: "$2a$10$PSxRcr1k2DMrUqrNHCPq2OS6dknAZm32fAaDV66hnPX8BEKxcEFKa", //test123
		Name:     "Test User",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "test@test.test").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.test",
		Password: "test123",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d status code, but got %d", http.StatusOK, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if resData.Token == "" {
		t.Fatal("Token is empty")
	}
}

func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test@test.test",
		Password: "wrong_test123",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected %d status code, but got %d", http.StatusUnauthorized, res.StatusCode)
	}
}
