package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type ICategoryService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.Category, *domain.AppError)
	Find(id uint) (*domain.Category, *domain.AppError)
	Save(dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError)
	Update(id uint, dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError)
	Delete(id uint) (*domain.Category, *domain.AppError)
}

type CategoryService struct {
	repository domain.ICategoryRepository
}

func (c *CategoryService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.Category, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return c.repository.List(criteria)
}

func (c *CategoryService) Find(id uint) (*domain.Category, *domain.AppError) {
	return c.repository.Find(id)
}

func (c *CategoryService) Save(dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	category := &domain.Category{
		Name:               dto.Name,
		AvailableFrom:      dto.AvailableFrom,
		AvailableUntil:     dto.AvailableUntil,
		AvailableFromHour:  dto.AvailableFromHour,
		AvailableUntilHour: dto.AvailableUntilHour,
		AvailableDays:      dto.AvailableDays,
	}
	return c.repository.Save(category)

}

func (c *CategoryService) Update(id uint, dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	category, err := c.repository.Find(id)
	if err != nil {
		return nil, err
	}
	category.Name = dto.Name
	category.AvailableFrom = dto.AvailableFrom
	category.AvailableUntil = dto.AvailableUntil
	category.AvailableFromHour = dto.AvailableFromHour
	category.AvailableUntilHour = dto.AvailableUntilHour
	category.AvailableDays = dto.AvailableDays
	return c.repository.Update(category)
}

func (c *CategoryService) Delete(id uint) (*domain.Category, *domain.AppError) {
	return c.repository.Delete(id)
}

func NewCategoryService(repository domain.ICategoryRepository) *CategoryService {
	return &CategoryService{repository}
}
