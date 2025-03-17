package orders

import (
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/user"
	"gorm.io/gorm"
)

type OrderItem struct {
	OrderId   uint `json:"orderId" gorm:"uniqueIndex:order_item"`
	Order     Order
	ProductId uint `json:"productId" gorm:"uniqueIndex:order_item"`
	Product   products.Product
	Count     int `json:"count" gorm:"default:1"`
	Price     int `json:"price" gorm:"default:1"`
}

type Order struct {
	gorm.Model
	UserId uint `json:"userId" gorm:"foreignKey:ID;"`
	User   user.User
	Items  []OrderItem `json:"items"`
	Total  int         `json:"total"`
}
