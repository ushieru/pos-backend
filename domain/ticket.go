package domain

import domain_criteria "github.com/ushieru/pos/domain/criteria"

type TicketStatus string

const (
	TicketOpen   TicketStatus = "open"
	TicketCancel TicketStatus = "cancel"
	TicketPaid   TicketStatus = "paid"
)

type Ticket struct {
	Model
	TicketStatus   TicketStatus    `json:"ticket_status"`
	Total          float64         `json:"total" gorm:"default:0"`
	AccountID      string          `json:"account_id"`
	Account        Account         `json:"account"`
	TicketProducts []TicketProduct `json:"ticket_products"`
}

type ITicketRepository interface {
	List(criteria *domain_criteria.Criteria) ([]Ticket, *AppError)
	Save(*Ticket) (*Ticket, *AppError)
	Find(id string) (*Ticket, *AppError)
	Delete(*Ticket) (*Ticket, *AppError)
	AddProduct(ticket *Ticket, product *Product) (*Ticket, *AppError)
	DeleteProduct(ticket *Ticket, product *Product) (*Ticket, *AppError)
	PayTicket(ticket *Ticket) (*Ticket, *AppError)
	OrderTicketProductsByTicket(*Ticket) (*Ticket, *AppError)
}
