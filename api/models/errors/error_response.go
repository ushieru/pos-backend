package models_errors

type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
	Details     string `json:"details"`
}

func NewErrorResponse(error string, description string, details string) ErrorResponse {
	return ErrorResponse{
		Error:       error,
		Description: description,
		Details:     details,
	}
}
