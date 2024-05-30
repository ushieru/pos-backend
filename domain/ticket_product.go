package domain

type TicketProduct struct {
	Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	IsEditable  *bool   `json:"is_editable"`
	ProductID   string  `json:"product_id"`
	TicketID    string  `json:"ticket_id"`
}
