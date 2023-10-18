package dto

type UpdateCategoryDTO struct {
	Name string `json:"name" validate:"required,min=3"`
}
