package dto

import "github.com/ushieru/pos/domain"

type UpdateTableRequest struct {
	Name string `json:"name"`
	PosX uint   `json:"pos_x"`
	PosY uint   `json:"pos_y"`
}

func (dto *UpdateTableRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	if dto.PosX < 1 || dto.PosY < 1 {
		return domain.NewValidationError("No se permiten pocisiones negativas")
	}
	return nil
}
