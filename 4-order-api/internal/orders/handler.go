package orders

import (
	"fmt"
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/middleware"
	"net/http"
)

type ordersHandler struct {
	*configs.Config
	*OrdersRepository
}

type OrdersHandlerDependencies struct {
	*configs.Config
	*OrdersRepository
}

func NewOrdersHandler(router *http.ServeMux, dependencies OrdersHandlerDependencies) {
	handler := &ordersHandler{
		Config:           dependencies.Config,
		OrdersRepository: dependencies.OrdersRepository,
	}
	router.Handle("POST /orders", middleware.IsAuthed(handler.createOrder(), dependencies.Config))
	router.Handle("GET /orders/{orderId}", middleware.IsAuthed(handler.getOrderById(), dependencies.Config))
	router.Handle("GET /orders", middleware.IsAuthed(handler.findUserOrders(), dependencies.Config))
}

func (handler *ordersHandler) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("createOrder")
		// Validate body

		// Create order

		// Response
	}
}

func (handler *ordersHandler) getOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getById")
		// Validate params

		// Find order

		// Response
	}
}

func (handler *ordersHandler) findUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("findUserOrders")
		// Validate params

		// Get userId by phone

		// Find orders by userId

		// Response
	}
}
