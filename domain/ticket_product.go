package domain

type TicketProduct struct {
	Product
	ProductId uint   `json:"product_id"`
	Quantity  uint16 `json:"quantity"`
	TicketID  uint   `json:"ticket_id"`
}
