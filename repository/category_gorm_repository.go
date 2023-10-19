package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	database *gorm.DB
}

func (r *CategoryGormRepository) List() ([]domain.Category, *domain.AppError) {
	var category []domain.Category
	r.database.Preload("Products").Find(&category)
	return category, nil
}

func (r *CategoryGormRepository) Save(category *domain.Category) (*domain.Category, *domain.AppError) {
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

func (r *CategoryGormRepository) Update(c *domain.Category) (*domain.Category, *domain.AppError) {
	category, err := r.Find(c.ID)
	if err != nil {
		return nil, err
	}
	category.Name = c.Name
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

func NewCategoryGormRepository(database *gorm.DB) *CategoryGormRepository {
	database.AutoMigrate(&domain.Category{})
	return &CategoryGormRepository{database}
}
