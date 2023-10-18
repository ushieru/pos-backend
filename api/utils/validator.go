package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/ushieru/pos/api/models/errors"
)

var validate = validator.New()

func ValidateStruct(model interface{}) *models_errors.ValidatorErrorResponse {
	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var validationError models_errors.ValidatorErrorResponse
			validationError.FailedField = err.StructNamespace()
			validationError.Tag = err.Tag()
			validationError.Value = err.Param()
			return &validationError
		}
	}
	return nil
}
