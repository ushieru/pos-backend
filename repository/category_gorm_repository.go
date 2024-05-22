package repository

import (
	"fmt"
	"time"

	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *CategoryGormRepository) seed() {
	category := new(domain.Category)
	r.database.First(category)
	if category.ID != 0 {
		return
	}
	for i := 1; i <= 5; i++ {
		r.database.Create(&domain.Category{
			Name:           fmt.Sprintf("Category %d", i),
			AvailableFrom:  time.Now(),
			AvailableUntil: time.Now().AddDate(1, 0, 0),
		})
	}
}

func (r *CategoryGormRepository) List(criteria *domain_criteria.Criteria) ([]domain.Category, *domain.AppError) {
	var category []domain.Category
	scopes := r.c.FiltersToScopes(criteria.Filters)
	tx := r.database.Preload("Products")
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

func (r *CategoryGormRepository) Find(id uint) (*domain.Category, *domain.AppError) {
	category := new(domain.Category)
	r.database.Preload("Products").First(category, id)
	if category.ID == 0 {
		return nil, domain.NewNotFoundError("Categoria no encontrada")
	}
	return category, nil
}

func (r *CategoryGormRepository) Update(category *domain.Category) (*domain.Category, *domain.AppError) {
	r.database.Save(category)
	return category, nil
}

func (r *CategoryGormRepository) Delete(id uint) (*domain.Category, *domain.AppError) {
	category, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(category)
	return category, nil
}

func NewCategoryGormRepository(database *gorm.DB) domain.ICategoryRepository {
	database.AutoMigrate(&domain.Category{})
	r := CategoryGormRepository{database: database}
	r.seed()
	return &r
}
