package dto

import "github.com/ushieru/pos/domain"

type CreateUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccountType string `json:"account_type"`
}

func (dto *CreateUserRequest) Validate() *domain.AppError {
	if dto.Name == "" {
		return domain.NewValidationError("No se permiten nombres vacios")
	}
	if dto.Email == "" {
		return domain.NewValidationError("No se permiten emails vacios")
	}
	if dto.Username == "" {
		return domain.NewValidationError("No se permiten usuarios vacios")
	}
	if len(dto.Password) < 5 {
		return domain.NewValidationError("No se permiten contraseñas menores a 5 caracteres")
	}
	if dto.AccountType != string(domain.Admin) &&
		dto.AccountType != string(domain.Cashier) &&
		dto.AccountType != string(domain.Waiter) {
		return domain.NewValidationError("AccountType permitidos: admin, cashier o waiter")
	}
	return nil
}
