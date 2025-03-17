package orders

import (
	"go-ps-adv-homework/configs"
	"go-ps-adv-homework/pkg/db"
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

func (repository *OrdersRepository) Create(order *Order) (*Order, error) {
	//err := repository.Database.DB.Transaction(func(tx *gorm.DB) error {
	//	// Создаем заказ
	//	if err := tx.Table("orders").Create(&order).Error; err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	// Устанавливаем OrderId для каждого элемента заказа
	//	for i := range order.Items {
	//		order.Items[i].OrderId = order.ID
	//	}
	//	// Создаем элементы заказа пачками
	//	if err := tx.CreateInBatches(order.Items, 100).Error; err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return nil, err
	//}
	//return order, nil
	result := repository.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (repository *OrdersRepository) GetById(id uint) (*Order, error) {
	var product Order
	result := repository.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *OrdersRepository) Find(from, to time.Time, userId uint, limit int, offset int) (*[]Order, error) {
	var orders []Order
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
		Model(&Order{}).
		Where("user_id = ?", userId).
		Where("created_at BETWEEN ? AND ?", from, to).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
