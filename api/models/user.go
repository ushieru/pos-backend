package models

type User struct {
	Model

	Name  string `json:"name"`
	Email string `json:"email"`

	Account Account `json:"account" validate:"dive"`
}
