package orders

import (
	"errors"
	"go-ps-adv-homework/internal/users"
	"go-ps-adv-homework/pkg/di"
	"time"
)

type OrdersService struct {
	UserRepository   *users.UserRepository
	OrdersRepository di.IOrdersRepository
}

type OrdersServiceDependencies struct {
	UserRepository   *users.UserRepository
	OrdersRepository di.IOrdersRepository
}

func NewOrdersService(dependencies OrdersServiceDependencies) *OrdersService {
	return &OrdersService{
		UserRepository:   dependencies.UserRepository,
		OrdersRepository: dependencies.OrdersRepository,
	}
}

func (service *OrdersService) CreateOrder(phone string, body di.CreateOrderRequest) (*di.Order, error) {
	// 1. Получаем пользователя по номеру телефона
	user, err := service.UserRepository.GetUserByPhone(phone)
	if err != nil {
		return nil, errors.New("пользователь не найден")
	}
	// 2. Подсчет общей суммы заказа и формирование списка товаров
	var total float64
	var items []di.OrderItem

	for _, item := range body.Items {
		total += item.Price * float64(item.Count)
		items = append(items, di.OrderItem{
			ProductID: item.ProductID,
			Count:     item.Count,
			Price:     item.Price,
		})
	}
	// 3. Создаем объект заказа
	order := &di.Order{
		UserID: user.ID,
		Total:  total,
		Items:  items,
	}
	// 4. Вызываем репозиторий для сохранения заказа
	createdOrder, err := service.OrdersRepository.Create(order)
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (service *OrdersService) GetUserOrders(phone string, from, to time.Time, limit, offset int) (*[]di.Order, int64, error) {
	user, err := service.UserRepository.GetUserByPhone(phone)
	if err != nil {
		return nil, 0, errors.New("пользователь не найден")
	}

	orders, err := service.OrdersRepository.Find(from, to, user.ID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := service.OrdersRepository.Count(from, to, user.ID)
	if err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

func (service *OrdersService) GetOrderByID(orderID uint) (*di.Order, error) {
	order, err := service.OrdersRepository.GetById(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
