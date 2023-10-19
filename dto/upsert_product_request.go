package dto

import "github.com/ushieru/pos/domain"

type UpsertProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (dto *UpsertProductRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	if dto.Description == "" {
		return domain.NewValidationError("No se permiten descriptiones vacias")
	}
	if dto.Price < 0 {
		return domain.NewValidationError("No se permiten precios negativos")
	}
	return nil
}
