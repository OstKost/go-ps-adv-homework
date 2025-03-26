package orders

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/di"
	"go-ps-adv-homework/pkg/middleware"
	"go-ps-adv-homework/pkg/request"
	"go-ps-adv-homework/pkg/response"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ordersHandler struct {
	*configs.Config
	OrdersService di.IOrdersService
}

type OrdersHandlerDependencies struct {
	*configs.Config
	OrdersService di.IOrdersService
}

func NewOrdersHandler(router *http.ServeMux, dependencies OrdersHandlerDependencies) {
	handler := &ordersHandler{
		Config:        dependencies.Config,
		OrdersService: dependencies.OrdersService,
	}
	router.Handle("POST /orders", middleware.IsAuthed(handler.createOrder(), dependencies.Config))
	router.Handle("GET /orders/{orderId}", middleware.IsAuthed(handler.getOrderById(), dependencies.Config))
	router.Handle("GET /orders", middleware.IsAuthed(handler.findUserOrders(), dependencies.Config))
}

func (handler *ordersHandler) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Context
		ctx := r.Context()
		phone, _ := ctx.Value(middleware.ContextPhoneKey).(string)
		// Validate
		body, err := request.HandleBody[di.CreateOrderRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Create order
		createdOrder, err := handler.OrdersService.CreateOrder(phone, *body)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, createdOrder, http.StatusCreated)
	}
}

func (handler *ordersHandler) getOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate id
		idString := r.PathValue("orderId")
		if strings.Trim(idString, " ") == "" {
			response.Json(w, "order ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Action
		order, err := handler.OrdersService.GetOrderByID(uint(id))
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, order, http.StatusOK)
	}
}

func (handler *ordersHandler) findUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		limit, offset := request.GetPaginationParams(query)
		fromString := query.Get("from")
		toString := query.Get("to")
		// Parse dates
		from, err := time.Parse("2006-01-02", fromString)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		to, err := time.Parse("2006-01-02", toString)
		if err != nil {
			log.Println(err.Error())
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Context
		ctx := r.Context()
		phone, _ := ctx.Value(middleware.ContextPhoneKey).(string)
		// Get orders
		orders, total, err := handler.OrdersService.GetUserOrders(phone, from, to, limit, offset)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
		}
		responseData := response.PreparePaginatedResponse[di.Order](*orders, total, offset, limit)
		response.Json(w, responseData, http.StatusOK)
	}
}
