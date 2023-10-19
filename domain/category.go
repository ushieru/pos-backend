package domain

type Category struct {
	Model
	Name     string    `json:"name"`
	Products []Product `gorm:"many2many:category_product;" json:"products"`
}

type ICategoryRepository interface {
	List() ([]Category, *AppError)
	Save(*Category) (*Category, *AppError)
	Find(id uint) (*Category, *AppError)
	Update(*Category) (*Category, *AppError)
	Delete(id uint) (*Category, *AppError)
}
