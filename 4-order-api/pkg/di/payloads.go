package di

type CreateOrderRequest struct {
	Items []OrderItem `json:"items"`
}
