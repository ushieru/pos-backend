package models

type Product struct {
	Model

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	Categories []Category `gorm:"many2many:category_product;" json:"categories"`
}
