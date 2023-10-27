package dto

import "github.com/ushieru/pos/domain"

type UpsertCategoryRequest struct {
	Name string `json:"name"`
}

func (dto *UpsertCategoryRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	return nil
}
