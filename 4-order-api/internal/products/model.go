package products

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name" gorm:"size:100; uniqueIndex"`
	Description string         `json:"description" gorm:"size:500"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
}

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required,min=3,max=50"`
	Description string   `json:"description" validate:"max=500"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name" validate:"min=3,max=50"`
	Description string   `json:"description" validate:"max=500"`
	Images      []string `json:"images"`
	Price       float64  `json:"price"`
}
