package products

import (
	"go-ps-adv-homework/pkg/db"
	"gorm.io/gorm/clause"
)

type ProductsRepository struct {
	Database *db.Db
}

func NewProductsRepository(database *db.Db) *ProductsRepository {
	return &ProductsRepository{
		Database: database,
	}
}

func (repository *ProductsRepository) Create(product *Product) (*Product, error) {
	result := repository.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repository *ProductsRepository) Update(product *Product) (*Product, error) {
	result := repository.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repository *ProductsRepository) Delete(id uint) error {
	result := repository.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *ProductsRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repository.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repository *ProductsRepository) Find(name string, limit int, offset int) (*[]Product, error) {
	var products []Product
	result := repository.Database.DB.
		Where("name LIKE ?", "%"+name+"%").
		Limit(limit).
		Offset(offset).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}

func (repository *ProductsRepository) Count(name string) (int64, error) {
	var count int64
	result := repository.Database.DB.
		Model(&Product{}).
		Where("name LIKE ?", "%"+name+"%").
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
