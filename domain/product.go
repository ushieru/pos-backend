package domain

import (
	"time"

	domain_criteria "github.com/ushieru/pos/domain/criteria"
)

type Product struct {
	Model
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	Price              float64    `json:"price"`
	AvailableFrom      time.Time  `json:"available_from"`
	AvailableUntil     time.Time  `json:"available_until"`
	AvailableDays      string     `json:"available_days"`
	AvailableFromHour  string     `json:"available_from_hour"`
	AvailableUntilHour string     `json:"available_until_hour"`
	Categories         []Category `json:"categories"           gorm:"many2many:category_product;"`
}

type IProductRepository interface {
	List(*domain_criteria.Criteria) ([]Product, *AppError)
	ListByCategoryId(id uint, criteria *domain_criteria.Criteria) ([]Product, *AppError)
	Save(*Product) (*Product, *AppError)
	Find(id uint) (*Product, *AppError)
	Update(*Product) (*Product, *AppError)
	Delete(id uint) (*Product, *AppError)
	AddCategory(productId, categoryId uint) (*Product, *AppError)
	DeleteCategory(productId, categoryId uint) (*Product, *AppError)
}
