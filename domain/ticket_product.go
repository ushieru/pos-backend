package domain

import domain_criteria "github.com/ushieru/pos/domain/criteria"

type TicketProductStatus string

const (
	Added         TicketProductStatus = "Added"
	Ordered       TicketProductStatus = "Ordered"
	InPreparation TicketProductStatus = "InPreparation"
	Prepared      TicketProductStatus = "Prepared"
)

type TicketProduct struct {
	Model
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Price       float64             `json:"price"`
	Status      TicketProductStatus `json:"status"`
	ProductID   string              `json:"product_id"`
	TicketID    string              `json:"ticket_id"`
}

type ITicketProductRepository interface {
	List(*domain_criteria.Criteria) ([]TicketProduct, *AppError)
	FindByAccountProductionCenters(account *Account) ([]TicketProduct, *AppError)
	Find(id string, criteria *domain_criteria.Criteria) (*TicketProduct, *AppError)
	Update(*TicketProduct) (*TicketProduct, *AppError)
}
