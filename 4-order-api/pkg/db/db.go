package db

import (
	"go-ps-adv-homework/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(config *configs.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DB.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
