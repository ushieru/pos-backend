package models

type TicketStatus string

const (
	TicketOpen  TicketStatus = "open"
	TicketClose              = "close"
)

type Ticket struct {
	Model

	TicketStatus TicketStatus `json:"ticket_status"`
	Total        float64      `json:"total" gorm:"default:0"`

	AccountID uint    `json:"account_id"`
	Account   Account `json:"account"`

	TicketProducts []TicketProduct `json:"ticket_products"`
}
