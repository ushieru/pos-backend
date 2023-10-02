package models

type AccountType string

const (
	Admin   AccountType = "admin"
	Cashier             = "cashier"
	Waiter              = "waiter"
)

type Account struct {
	Model

	Username    string      `json:"username" validate:"required,min=3"`
	Password    string      `json:"-" validate:"required,min=3"`
	IsActive    *bool       `json:"is_active" gorm:"default:true"`
	AccountType AccountType `json:"account_type" gorm:"default:waiter"`
	UserID      uint        `json:"user_id"`
}
