package carts

import (
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/users"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId uint       `json:"userId"`
	User   users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CartItem struct {
	ID        uint             `json:"id" gorm:"primaryKey,autoIncrement"`
	CartId    uint             `json:"cartId"`
	Cart      Cart             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductId uint             `json:"productId"`
	Product   products.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Count     int              `json:"count" gorm:"default:1"`
}
