package domain

type TicketProduct struct {
	Model
	Product  Product `json:"product" gorm:"embedded"`
	Quantity uint16  `json:"quantity"`
	TicketID uint    `json:"ticket_id"`
}
