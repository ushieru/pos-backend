package dto

import (
	"time"

	"github.com/ushieru/pos/domain"
)

type UpsertCategoryRequest struct {
	Name               string    `json:"name"`
	AvailableFrom      time.Time `json:"available_from"`
	AvailableUntil     time.Time `json:"available_until"`
	AvailableFromHour  string    `json:"available_from_hour"`
	AvailableUntilHour string    `json:"available_until_hour"`
	AvailableDays      string    `json:"available_days"`
}

func (dto *UpsertCategoryRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	if dto.AvailableUntil.Before(dto.AvailableFrom) {
		return domain.NewValidationError("No se permite una fecha anterior a la fecha de disponibilidad")
	}
	// TODO: Validate -> AvailableFromHour, AvailableUntilHour, AvailableDays
	return nil
}
