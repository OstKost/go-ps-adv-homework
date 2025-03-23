package orders

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/db"
	"go-ps-adv-homework/pkg/di"
	"time"
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

func (repository *OrdersRepository) Create(order *di.Order) (*di.Order, error) {
	result := repository.Database.DB.Create(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repository *OrdersRepository) GetById(id uint) (*di.Order, error) {
	var product di.Order
	result := repository.Database.DB.
		Preload("User").Preload("Items.Product").
		First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *OrdersRepository) Find(from, to time.Time, userId uint, limit int, offset int) (*[]di.Order, error) {
	var orders []di.Order
	result := repository.Database.DB.
		Where("user_id = ?", userId).
		Where("created_at BETWEEN ? AND ?", from, to).
		Limit(limit).
		Offset(offset).
		Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}

func (repository *OrdersRepository) Count(from, to time.Time, userId uint) (int64, error) {
	var count int64
	result := repository.Database.DB.
		Model(&di.Order{}).
		Where("user_id = ?", userId).
		Where("created_at BETWEEN ? AND ?", from, to).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
