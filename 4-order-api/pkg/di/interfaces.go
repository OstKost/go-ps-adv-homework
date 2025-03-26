package di

import (
	"time"
)

type IOrdersRepository interface {
	Create(order *Order) (*Order, error)
	GetById(id uint) (*Order, error)
	Find(from, to time.Time, userId uint, limit int, offset int) (*[]Order, error)
	Count(from, to time.Time, userId uint) (int64, error)
}

type IOrdersService interface {
	CreateOrder(phone string, body CreateOrderRequest) (*Order, error)
	GetUserOrders(phone string, from, to time.Time, limit, offset int) (*[]Order, int64, error)
	GetOrderByID(orderID uint) (*Order, error)
}
