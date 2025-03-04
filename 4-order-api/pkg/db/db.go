package db

import (
	"go-ps-adv-homework/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func Connect(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.DB.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{DB: db}
}
