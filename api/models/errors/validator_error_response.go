package models_errors

import "fmt"

type ValidatorErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func (v ValidatorErrorResponse) ToString() string {
	return fmt.Sprintf("%s %s %s", v.FailedField, v.Tag, v.Value)
}
