package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type ProductGormRepository struct {
	database *gorm.DB
}

func (r *ProductGormRepository) List() ([]domain.Product, *domain.AppError) {
	var products []domain.Product
	r.database.Preload("Categories").Find(&products)
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

func (r *ProductGormRepository) AddCategory(productId, categoryId uint) (*domain.Product, *domain.AppError) {
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

func NewProductGormRepository(database *gorm.DB) *ProductGormRepository {
	database.AutoMigrate(&domain.Product{})
	return &ProductGormRepository{database}
}