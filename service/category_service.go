package service

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
)

type ICategoryService interface {
	List() ([]domain.Category, *domain.AppError)
	Find(id uint) (*domain.Category, *domain.AppError)
	Save(dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError)
	Update(id uint, dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError)
	Delete(id uint) (*domain.Category, *domain.AppError)
}

type CategoryService struct {
	repository domain.ICategoryRepository
}

func (c *CategoryService) List() ([]domain.Category, *domain.AppError) {
	return c.repository.List()
}

func (c *CategoryService) Find(id uint) (*domain.Category, *domain.AppError) {
	return c.repository.Find(id)
}

func (c *CategoryService) Save(dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	category := &domain.Category{
		Name: dto.Name,
	}
	return c.repository.Save(category)

}

func (c *CategoryService) Update(id uint, dto *dto.UpsertCategoryRequest) (*domain.Category, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	category := &domain.Category{
		Model: domain.Model{ID: id},
		Name:  dto.Name,
	}
	return c.repository.Update(category)
}

func (c *CategoryService) Delete(id uint) (*domain.Category, *domain.AppError) {
	return c.repository.Delete(id)
}

func NewCategoryService(repository domain.ICategoryRepository) *CategoryService {
	return &CategoryService{repository}
}
