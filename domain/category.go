package domain

import (
	"time"

	domain_criteria "github.com/ushieru/pos/domain/criteria"
)

type Category struct {
	Model
	Name               string    `json:"name"`
	AvailableFrom      time.Time `json:"available_from"`
	AvailableUntil     time.Time `json:"available_until"`
	AvailableFromHour  string    `json:"available_from_hour"`
	AvailableUntilHour string    `json:"available_until_hour"`
	AvailableDays      string    `json:"available_days"`
	Products           []Product `gorm:"many2many:category_product;" json:"products"`
}

type ICategoryRepository interface {
	List(*domain_criteria.Criteria, bool) ([]Category, *AppError)
	Save(*Category) (*Category, *AppError)
	Find(id uint) (*Category, *AppError)
	Update(*Category) (*Category, *AppError)
	Delete(id uint) (*Category, *AppError)
}
