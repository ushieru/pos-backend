package models

type Category struct {
	Model

	Name string `json:"name"`

	Products []Product `gorm:"many2many:category_product;" json:"products"`
}
