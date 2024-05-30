package domain

type AccountType string

const (
	Admin   AccountType = "admin"
	Cashier AccountType = "cashier"
	Waiter  AccountType = "waiter"
)

type Account struct {
	Model
	Username    string      `json:"username"`
	Password    string      `json:"-"`
	IsActive    *bool       `json:"is_active"`
	AccountType AccountType `json:"account_type"`
	UserID      string      `json:"user_id"`
}
