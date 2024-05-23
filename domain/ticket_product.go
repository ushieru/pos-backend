package domain

type TicketProduct struct {
	Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsEditable  *bool   `json:"is_editable"`
	ProductID   uint    `json:"product_id"`
	TicketID    uint    `json:"ticket_id"`
}
