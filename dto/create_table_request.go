package dto

import "github.com/ushieru/pos/domain"

type CreateTableRequest struct {
	Name string `json:"name"`
}

func (dto *CreateTableRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	return nil
}
