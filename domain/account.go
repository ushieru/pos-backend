package domain

type AccountType string

const (
	Admin     AccountType = "admin"
	Cashier   AccountType = "cashier"
	Waiter    AccountType = "waiter"
	Cook      AccountType = "cook"
	Bartender AccountType = "bartender"
)

type Account struct {
	Model
	Username           string      `json:"username"`
	Password           string      `json:"-"`
	IsActive           *bool       `json:"is_active"`
	AccountType        AccountType `json:"account_type"`
	UserID             string      `json:"user_id"`
	ProductionCenterID string      `json:"-"`
}

type IAccountRepository interface {
	Find(id string) (*Account, *AppError)
}
