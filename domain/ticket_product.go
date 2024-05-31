package domain

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
