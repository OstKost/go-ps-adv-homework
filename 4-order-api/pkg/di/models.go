package di

import (
	"go-ps-adv-homework/internal/products"
	"go-ps-adv-homework/internal/users"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID uint
	User   users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total  int        `json:"total"`
	Items  []OrderItem
}

type OrderItem struct {
	ID        uint             `json:"id" gorm:"primaryKey,autoIncrement"`
	OrderID   uint             `json:"orderId" gorm:"uniqueIndex:idx_order_item_order_id_product_id"` //gorm:"uniqueIndex:idx_order_item_order_id_product_id"
	Order     Order            `gorm:"PRELOAD:false,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID uint             `json:"productId" gorm:"uniqueIndex:idx_order_item_order_id_product_id"` //gorm:"uniqueIndex:idx_order_item_order_id_product_id"
	Product   products.Product `gorm:"PRELOAD:false,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Count     int              `json:"count"`
	Price     int              `json:"price"`
}
