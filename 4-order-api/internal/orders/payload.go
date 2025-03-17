package orders

type CreateNewOrder struct {
	ID     uint                 `json:"id"`
	UserId uint                 `json:"userId"`
	Items  []CreateNewOrderItem `json:"items"`
	Total  int                  `json:"total"`
}

type CreateNewOrderItem struct {
	OrderId   uint `json:"orderId"`
	ProductId uint `json:"productId"`
	Count     int  `json:"count"`
	Price     int  `json:"price"`
}

type CreateOrderRequest struct {
	Items []CreateNewOrderItem `json:"items"`
}
