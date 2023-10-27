package domain

type AccountType string

const (
	Admin   AccountType = "admin"
	Cashier             = "cashier"
	Waiter              = "waiter"
)

type Account struct {
	Model
	Username    string      `json:"username"`
	Password    string      `json:"-"`
	IsActive    *bool       `json:"is_active"`
	AccountType AccountType `json:"account_type"`
	UserID      uint        `json:"user_id"`
}
