package domain

import domain_criteria "github.com/ushieru/pos/domain/criteria"

type ProductionCenter struct {
	Model
	Name       string     `json:"name"`
	Accounts   []Account  `gorm:"many2many:account_production_center;" json:"accounts"`
	Categories []Category `json:"categories"`
}

type IProductionCenterRepository interface {
	List(*domain_criteria.Criteria) ([]ProductionCenter, *AppError)
	Save(*ProductionCenter) (*ProductionCenter, *AppError)
	Find(id string) (*ProductionCenter, *AppError)
	Update(*ProductionCenter) (*ProductionCenter, *AppError)
	AddAccount(*ProductionCenter, *Account) (*ProductionCenter, *AppError)
	DeleteAccount(*ProductionCenter, *Account) (*ProductionCenter, *AppError)
	AddCategory(*ProductionCenter, *Category) (*ProductionCenter, *AppError)
	DeleteCategory(*ProductionCenter, *Category) (*ProductionCenter, *AppError)
	Delete(category *ProductionCenter) (*ProductionCenter, *AppError)
}
