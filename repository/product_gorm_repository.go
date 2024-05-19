package repository

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type ProductGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *ProductGormRepository) List(
	criteria *domain_criteria.Criteria,
) ([]domain.Product, *domain.AppError) {
	var products []domain.Product
	scopes := r.c.FiltersToScopes(criteria.Filters)
	tx := r.database.Preload("Categories")
	if len(scopes) > 0 {
		tx.Scopes(scopes...)
	}
	tx.Find(&products)
	return products, nil
}

func (r *ProductGormRepository) ListByCategoryId(
	id uint,
	criteria *domain_criteria.Criteria,
) ([]domain.Product, *domain.AppError) {
	var products []domain.Product
	scopes := r.c.FiltersToScopes(criteria.Filters)
	stm := r.database.Debug().
		Preload("Categories").
		Where("id IN (SELECT product_id FROM category_product WHERE category_id = ?)", id)
	if len(scopes) > 0 {
		stm.Scopes(scopes...)
	}
	stm.Find(&products)
	return products, nil
}

func (r *ProductGormRepository) Save(product *domain.Product) (*domain.Product, *domain.AppError) {
	result := r.database.Save(product)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear producto")
	}
	return product, nil
}

func (r *ProductGormRepository) Find(id uint) (*domain.Product, *domain.AppError) {
	product := new(domain.Product)
	r.database.Preload("Categories").First(product, id)
	if product.ID == 0 {
		return nil, domain.NewNotFoundError("Producto no encontrado")
	}
	return product, nil
}

func (r *ProductGormRepository) Update(p *domain.Product) (*domain.Product, *domain.AppError) {
	product, err := r.Find(p.ID)
	if err != nil {
		return nil, err
	}
	product.Name = p.Name
	product.Description = p.Description
	product.Price = p.Price
	product.AvailableFrom = p.AvailableFrom
	product.AvailableUntil = p.AvailableUntil
	r.database.Save(product)
	return product, nil
}

func (r *ProductGormRepository) Delete(id uint) (*domain.Product, *domain.AppError) {
	product, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(product)
	return product, nil
}

func (r *ProductGormRepository) AddCategory(
	productId, categoryId uint,
) (*domain.Product, *domain.AppError) {
	product, err := r.Find(productId)
	if err != nil {
		return nil, err
	}
	category := new(domain.Category)
	r.database.First(category, categoryId)
	if category.ID == 0 {
		return nil, domain.NewNotFoundError("Categoria no encontrada")
	}
	if err := r.database.Model(product).Association("Categories").Append(category); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible agregar la categoria al producto")
	}
	return product, nil
}

func (r *ProductGormRepository) DeleteCategory(
	productId, categoryId uint,
) (*domain.Product, *domain.AppError) {
	product, err := r.Find(productId)
	if err != nil {
		return nil, err
	}
	category := new(domain.Category)
	r.database.First(category, categoryId)
	if category.ID == 0 {
		return nil, domain.NewNotFoundError("Categoria no encontrada")
	}
	if err := r.database.Model(product).Association("Categories").Delete(category); err != nil {
		return nil, domain.NewUnexpectedError("No fue posible eliminar la categoria del producto")
	}
	return product, nil
}

func NewProductGormRepository(database *gorm.DB) domain.IProductRepository {
	database.AutoMigrate(&domain.Product{})
	return &ProductGormRepository{database: database}
}
