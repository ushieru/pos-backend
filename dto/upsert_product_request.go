package dto

import (
	"strconv"
	"strings"
	"time"

	"github.com/ushieru/pos/domain"
)

type UpsertProductRequest struct {
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Price              float64   `json:"price"`
	AvailableFrom      time.Time `json:"available_from" example:"2023-12-15T21:54:42.123Z"`
	AvailableUntil     time.Time `json:"available_until" example:"2023-12-18T21:54:42.123Z"`
	AvailableFromHour  string    `json:"available_from_hour" example:"00:00"`
	AvailableUntilHour string    `json:"available_until_hour" example:"00:00"`
	AvailableDays      string    `json:"available_days" example:"0,1,2,3,4,5"`
}

func (dto *UpsertProductRequest) ToProduct() *domain.Product {
	return &domain.Product{
		Name:               dto.Name,
		Description:        dto.Description,
		Price:              dto.Price,
		AvailableFrom:      dto.AvailableFrom,
		AvailableUntil:     dto.AvailableUntil,
		AvailableFromHour:  dto.AvailableFromHour,
		AvailableUntilHour: dto.AvailableUntilHour,
		AvailableDays:      dto.AvailableDays,
	}
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
	if err := validateHour(dto.AvailableFromHour); err != nil {
		return err
	}
	if err := validateHour(dto.AvailableUntilHour); err != nil {
		return err
	}
	availableFromHour := time.Now()
	availableFromHourSplit := strings.Split(dto.AvailableFromHour, ":")
	availableFromHourHour, _ := strconv.Atoi(availableFromHourSplit[0])
	availableFromHourMinute, _ := strconv.Atoi(availableFromHourSplit[1])
	availableUntilHour := time.Now()
	availableUntilHourSplit := strings.Split(dto.AvailableUntilHour, ":")
	availableUntilHourHour, _ := strconv.Atoi(availableUntilHourSplit[0])
	availableUntilHourMinute, _ := strconv.Atoi(availableUntilHourSplit[1])
	if time.Date(availableFromHour.Year(), availableFromHour.Month(), availableFromHour.Day(), availableFromHourHour, availableFromHourMinute, 0, 0, availableFromHour.Location()).
		After(time.Date(availableUntilHour.Year(), availableUntilHour.Month(), availableUntilHour.Day(), availableUntilHourHour, availableUntilHourMinute, 0, 0, availableUntilHour.Location())) {
		return domain.NewValidationError("No se permite una hora anterior a la hora de disponibilidad")
	}
	if dto.AvailableDays != "" {
		days := strings.Split(dto.AvailableDays, ",")
		for _, day := range days {
			intDay, err := strconv.Atoi(day)
			if err != nil {
				return domain.NewValidationError("Dias disponibles incorrectos")
			}
			if intDay < 0 || intDay > 6 {
				return domain.NewValidationError("Dias disponibles incorrectos")
			}
		}
	}
	return nil
}
