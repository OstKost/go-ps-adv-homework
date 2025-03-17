package orders

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/internal/users"
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
	*OrdersRepository
	*users.UserRepository
}

type OrdersHandlerDependencies struct {
	*configs.Config
	*OrdersRepository
	*users.UserRepository
}

func NewOrdersHandler(router *http.ServeMux, dependencies OrdersHandlerDependencies) {
	handler := &ordersHandler{
		Config:           dependencies.Config,
		OrdersRepository: dependencies.OrdersRepository,
		UserRepository:   dependencies.UserRepository,
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
		body, err := request.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			response.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Prepare data
		user, err := handler.UserRepository.GetUserByPhone(phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}
		order := NewOrder(user.ID, body.Items)
		//var items []OrderItem
		//for _, item := range body.Items {
		//	items = append(items, OrderItem{
		//		ProductId: item.ProductId,
		//		Count:     item.Count,
		//		Price:     item.Price,
		//	})
		//}
		//var total int
		//for _, item := range items {
		//	total += item.Price * item.Count
		//}
		//order := &Order{
		//	UserId: user.ID,
		//	Items:  items,
		//	Total:  total,
		//}
		// Action
		createdOrder, err := handler.OrdersRepository.Create(order)
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
		order, err := handler.OrdersRepository.GetById(uint(id))
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
		user, err := handler.UserRepository.GetUserByPhone(phone)
		if err != nil {
			response.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		orders, err := handler.OrdersRepository.Find(from, to, user.ID, limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		total, err := handler.OrdersRepository.Count(from, to, user.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responseData := response.PreparePaginatedResponse[Order](*orders, total, offset, limit)
		response.Json(w, responseData, http.StatusOK)
	}
}
