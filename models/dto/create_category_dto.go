package dto

type CreateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=3"`
}
