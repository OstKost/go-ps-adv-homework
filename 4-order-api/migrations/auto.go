package main

import (
	"github.com/joho/godotenv"
	"go-ps-adv-homework/internal/carts"
	"go-ps-adv-homework/internal/orders"
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/sessions"
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

	entities := []interface{}{
		&user.User{},
		&products.Product{},
		&carts.Cart{},
		&carts.CartItem{},
		&sessions.Session{},
		&orders.OrderItem{},
		&orders.Order{},
	}

	for _, entity := range entities {
		err := db.AutoMigrate(entity)
		if err != nil {
			panic(err)
		}
	}
}
