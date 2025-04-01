package main

import (
	"github.com/joho/godotenv"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/sessions"
	"go-ps-adv-homework/internal/users"
	"go-ps-adv-homework/pkg/di"
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

	entities := []interface{}{
		&users.User{},
		&sessions.Session{},
		&di.Order{},
		&di.OrderItem{},
		&products.Product{},
		&carts.CartItem{},
		&carts.Cart{},
	}

	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			panic(err)
		}
	}
}
