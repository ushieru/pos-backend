package repository

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type ProductionCenterGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *ProductionCenterGormRepository) List(criteria *domain_criteria.Criteria) ([]domain.ProductionCenter, *domain.AppError) {
	var productionCenter []domain.ProductionCenter
	scopes := r.c.FiltersToScopes(criteria.Filters)
	tx := r.database.Model(&domain.ProductionCenter{})
	if len(scopes) > 0 {
		tx.Scopes(scopes...)
	}
	tx.Find(&productionCenter)
	return productionCenter, nil
}

func (r *ProductionCenterGormRepository) Save(productionCenter *domain.ProductionCenter) (*domain.ProductionCenter, *domain.AppError) {
	result := r.database.Save(productionCenter)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear Centro de produccion")
	}
	return productionCenter, nil
}

func (r *ProductionCenterGormRepository) Find(id string) (*domain.ProductionCenter, *domain.AppError) {
	productionCenter := new(domain.ProductionCenter)
	r.database.
	Preload("Accounts").
	Preload("Categories").
	First(productionCenter, "id = ?", id)
	if productionCenter.ID == "" {
		return nil, domain.NewNotFoundError("Centro de produccion no encontrada")
	}
	return productionCenter, nil
}

func (r *ProductionCenterGormRepository) Update(productionCenter *domain.ProductionCenter) (*domain.ProductionCenter, *domain.AppError) {
	r.database.Save(productionCenter)
	return productionCenter, nil
}

func (r *ProductionCenterGormRepository) Delete(productionCenter *domain.ProductionCenter) (*domain.ProductionCenter, *domain.AppError) {
	r.database.Delete(productionCenter)
	return productionCenter, nil
}

func (r *ProductionCenterGormRepository) AddAccount(pc *domain.ProductionCenter, a *domain.Account) (*domain.ProductionCenter, *domain.AppError) {
	if err := r.database.Model(&pc).Association("Accounts").Append(a); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible realizar esta accion")
	}
	return pc, nil
}

func (r *ProductionCenterGormRepository) DeleteAccount(pc *domain.ProductionCenter, a *domain.Account) (*domain.ProductionCenter, *domain.AppError) {
	if err := r.database.Model(&pc).Association("Accounts").Delete(a); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible realizar esta accion")
	}
	return pc, nil

}

func (r *ProductionCenterGormRepository) AddCategory(pc *domain.ProductionCenter, c *domain.Category) (*domain.ProductionCenter, *domain.AppError) {
	if err := r.database.Model(&pc).Association("Categories").Append(c); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible realizar esta accion")
	}
	return pc, nil

}

func (r *ProductionCenterGormRepository) DeleteCategory(pc *domain.ProductionCenter, c *domain.Category) (*domain.ProductionCenter, *domain.AppError) {
	if err := r.database.Model(&pc).Association("Categories").Delete(c); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible realizar esta accion")
	}
	return pc, nil

}

func NewProductionCenterRepository(database *gorm.DB) domain.IProductionCenterRepository {
	database.AutoMigrate(&domain.ProductionCenter{})
	r := ProductionCenterGormRepository{database: database}
	return &r
}
