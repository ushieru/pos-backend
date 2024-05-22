package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type IProductService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.Product, *domain.AppError)
	ListByCategoryId(
		id uint,
		dto *dto.SearchCriteriaQueryRequest,
	) ([]domain.Product, *domain.AppError)
	Find(id uint) (*domain.Product, *domain.AppError)
	Save(dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError)
	Update(id uint, dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError)
	Delete(id uint) (*domain.Product, *domain.AppError)
	AddCategory(productId, categoryId uint) (*domain.Product, *domain.AppError)
	DeleteCategory(productId, categoryId uint) (*domain.Product, *domain.AppError)
}

type ProductService struct {
	repository domain.IProductRepository
}

func (s *ProductService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.Product, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.repository.List(criteria)
}

func (s *ProductService) ListByCategoryId(
	id uint,
	dto *dto.SearchCriteriaQueryRequest,
) ([]domain.Product, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.repository.ListByCategoryId(id, criteria)
}

func (s *ProductService) Find(id uint) (*domain.Product, *domain.AppError) {
	return s.repository.Find(id)
}

func (s *ProductService) Save(dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	product := &domain.Product{
		Name:           dto.Name,
		Description:    dto.Description,
		Price:          dto.Price,
		AvailableFrom:  dto.AvailableFrom,
		AvailableUntil: dto.AvailableUntil,
	}
	return s.repository.Save(product)
}

func (s *ProductService) Update(
	id uint,
	dto *dto.UpsertProductRequest,
) (*domain.Product, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	product := &domain.Product{
		Model:          domain.Model{ID: id},
		Name:           dto.Name,
		Description:    dto.Description,
		Price:          dto.Price,
		AvailableFrom:  dto.AvailableFrom,
		AvailableUntil: dto.AvailableUntil,
	}
	return s.repository.Update(product)
}

func (s *ProductService) Delete(id uint) (*domain.Product, *domain.AppError) {
	return s.repository.Delete(id)
}

func (s *ProductService) AddCategory(
	productId, categoryId uint,
) (*domain.Product, *domain.AppError) {
	return s.repository.AddCategory(productId, categoryId)
}

func (s *ProductService) DeleteCategory(
	productId, categoryId uint,
) (*domain.Product, *domain.AppError) {
	return s.repository.DeleteCategory(productId, categoryId)
}

func NewProductService(repository domain.IProductRepository) *ProductService {
	return &ProductService{repository}
}
