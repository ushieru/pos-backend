package dto

import "github.com/ushieru/pos/domain"

type CreateProductionCenterRequest struct {
	Name string
}

func (dto *CreateProductionCenterRequest) Validate() *domain.AppError {
	return nil
}
