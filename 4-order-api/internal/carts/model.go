package carts

import (
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/users"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint       `json:"userId" gorm:"foreignKey:ID"`
	User   users.User `json:"users" gorm:"foreignKey:UserId"`
}

type CartItem struct {
	CartId    uint             `json:"cartId" gorm:"foreignKey:ID; uniqueIndex:cart_item"`
	Cart      Cart             `json:"cart" gorm:"foreignKey:CartId"`
	ProductId uint             `json:"productId" gorm:"foreignKey:ID; uniqueIndex:cart_item"`
	Product   products.Product `json:"product" gorm:"foreignKey:ProductId"`
	Count     int              `json:"count" gorm:"default:1"`
}
