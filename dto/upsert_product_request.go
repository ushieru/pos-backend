package dto

import (
	"time"

	"github.com/ushieru/pos/domain"
)

type UpsertProductRequest struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	AvailableFrom  time.Time `json:"available_from" example:"2023-12-15T21:54:42.123Z"`
	AvailableUntil time.Time `json:"available_until" example:"2023-12-18T21:54:42.123Z"`
	Price          float64   `json:"price"`
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
	if dto.AvailableUntil.Before(dto.AvailableFrom) {
		return domain.NewValidationError("No se permite una fecha anterior a la fecha de disponibilidad")
	}
	return nil
}
