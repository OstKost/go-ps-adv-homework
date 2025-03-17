package orders

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/db"
)

type OrdersRepository struct {
	*configs.Config
	Database *db.Db
}

func NewOrdersRepository(database *db.Db) *OrdersRepository {
	return &OrdersRepository{
		Database: database,
	}
}

func (repository *OrdersRepository) Create(order *Order) (*Order, error) {
	result := repository.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}
