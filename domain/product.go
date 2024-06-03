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
	Categories         []Category `json:"categories" gorm:"many2many:category_product;"`
}

type IProductRepository interface {
	List(*domain_criteria.Criteria, bool) ([]Product, *AppError)
	ListByCategoryId(id string, criteria *domain_criteria.Criteria) ([]Product, *AppError)
	Save(*Product) (*Product, *AppError)
	Find(id string, criteria *domain_criteria.Criteria) (*Product, *AppError)
	Update(*Product) (*Product, *AppError)
	Delete(*Product) (*Product, *AppError)
	AddCategory(product *Product, category *Category) (*Product, *AppError)
	DeleteCategory(product *Product, category *Category) (*Product, *AppError)
}
