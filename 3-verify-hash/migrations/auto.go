package main

import (
	"github.com/joho/godotenv"
	"go-ps-adv-homework/internal/link"
	"go-ps-adv-homework/internal/stat"
	"go-ps-adv-homework/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&link.Link{},
		&stat.Stat{},
		&user.User{})
	if err != nil {
		panic(err)
	}

}
