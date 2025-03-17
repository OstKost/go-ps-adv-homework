package orders

import (
	"gorm.io/gorm"
)

func NewOrder(userId uint, items []CreateNewOrderItem) *Order {
	var total int
	for _, item := range items {
		total += item.Price * item.Count
	}
	var orderItems []OrderItem
	for _, item := range items {
		orderItems = append(orderItems, OrderItem{
			ProductId: item.ProductId,
			Count:     item.Count,
			Price:     item.Price,
		})
	}
	order := &Order{
		UserId: userId,
		Items:  orderItems,
		Total:  total,
	}
	return order
}

type Order struct {
	gorm.Model
	UserId uint        `json:"userId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items  []OrderItem `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total  int         `json:"total"`
}

type OrderItem struct {
	OrderId   uint `json:"orderId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductId uint `json:"productId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Count     int  `json:"count"`
	Price     int  `json:"price"`
}
