package repository

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *CategoryGormRepository) List(criteria *domain_criteria.Criteria, withProducts bool) ([]domain.Category, *domain.AppError) {
	var category []domain.Category
	scopes := r.c.FiltersToScopes(criteria.Filters)
	tx := r.database.Model(&domain.Category{})
	if withProducts {
		tx.Preload("Products")
	}
	if len(scopes) > 0 {
		tx.Scopes(scopes...)
	}
	tx.Find(&category)
	return category, nil
}

func (r *CategoryGormRepository) Save(
	category *domain.Category,
) (*domain.Category, *domain.AppError) {
	result := r.database.Save(category)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear categoria")
	}
	return category, nil
}

func (r *CategoryGormRepository) Find(id string) (*domain.Category, *domain.AppError) {
	category := new(domain.Category)
	r.database.Preload("Products").First(category, "id = ?", id)
	if category.ID == "" {
		return nil, domain.NewNotFoundError("Categoria no encontrada")
	}
	return category, nil
}

func (r *CategoryGormRepository) Update(category *domain.Category) (*domain.Category, *domain.AppError) {
	r.database.Save(category)
	return category, nil
}

func (r *CategoryGormRepository) Delete(category *domain.Category) (*domain.Category, *domain.AppError) {
	r.database.Delete(category)
	return category, nil
}

func NewCategoryGormRepository(database *gorm.DB) domain.ICategoryRepository {
	database.AutoMigrate(&domain.Category{})
	r := CategoryGormRepository{database: database}
	return &r
}
