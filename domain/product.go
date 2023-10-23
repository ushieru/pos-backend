package domain

type Product struct {
	Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Categories  []Category `gorm:"many2many:category_product;" json:"categories"`
}

type IProductRepository interface {
	List() ([]Product, *AppError)
	Save(*Product) (*Product, *AppError)
	Find(id uint) (*Product, *AppError)
	Update(*Product) (*Product, *AppError)
	Delete(id uint) (*Product, *AppError)
	AddCategory(productId, categoryId uint) (*Product, *AppError)
	DeleteCategory(productId, categoryId uint) (*Product, *AppError)
}
