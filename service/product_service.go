package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type IProductService interface {
	List(*dto.SearchCriteriaQueryRequest, bool) ([]domain.Product, *domain.AppError)
	ListByCategoryId(id string, dto *dto.SearchCriteriaQueryRequest) ([]domain.Product, *domain.AppError)
	Find(id string) (*domain.Product, *domain.AppError)
	Save(dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError)
	Update(id string, dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError)
	Delete(id string) (*domain.Product, *domain.AppError)
	AddCategory(productId, categoryId string) (*domain.Product, *domain.AppError)
	DeleteCategory(productId, categoryId string) (*domain.Product, *domain.AppError)
}

type ProductService struct {
	productRepository  domain.IProductRepository
	categoryRepository domain.ICategoryRepository
}

func (s *ProductService) List(dto *dto.SearchCriteriaQueryRequest, withCategories bool) ([]domain.Product, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.productRepository.List(criteria, withCategories)
}

func (s *ProductService) ListByCategoryId(id string, dto *dto.SearchCriteriaQueryRequest) ([]domain.Product, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.productRepository.ListByCategoryId(id, criteria)
}

func (s *ProductService) Find(id string) (*domain.Product, *domain.AppError) {
	return s.productRepository.Find(id, nil)
}

func (s *ProductService) Save(dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	product := dto.ToProduct()
	return s.productRepository.Save(product)
}

func (s *ProductService) Update(id string, dto *dto.UpsertProductRequest) (*domain.Product, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	if _, err := s.Find(id); err != nil {
		return nil, err
	}
	product := dto.ToProduct()
	product.ID = id
	return s.productRepository.Update(product)
}

func (s *ProductService) Delete(id string) (*domain.Product, *domain.AppError) {
	product, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	return s.productRepository.Delete(product)
}

func (s *ProductService) AddCategory(productId, categoryId string) (*domain.Product, *domain.AppError) {
	product, err := s.productRepository.Find(productId, nil)
	if err != nil {
		return nil, err
	}
	category, err := s.categoryRepository.Find(categoryId)
	if err != nil {
		return nil, err
	}
	return s.productRepository.AddCategory(product, category)
}

func (s *ProductService) DeleteCategory(productId, categoryId string) (*domain.Product, *domain.AppError) {
	product, err := s.productRepository.Find(productId, nil)
	if err != nil {
		return nil, err
	}
	category, err := s.categoryRepository.Find(categoryId)
	if err != nil {
		return nil, err
	}
	return s.productRepository.DeleteCategory(product, category)
}

func NewProductService(productRepository domain.IProductRepository, categoryRepository domain.ICategoryRepository) *ProductService {
	return &ProductService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}
