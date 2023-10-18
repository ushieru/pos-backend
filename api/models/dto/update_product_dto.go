package dto

type UpdateProductDTO struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description string  `json:"description" validate:"required,min=3"`
	Price       float64 `json:"price" validate:"required,gte=0"`
}

// type UpdateProductDTO struct {
// Name        string  `json:"name" validate:"omitempty,min=3"`
// Description string  `json:"description" validate:"omitempty,min=3"`
// Price       float64 `json:"price" validate:"omitempty,gte=0"`
// }
