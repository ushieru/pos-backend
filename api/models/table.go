package models

type Table struct {
	Model

	Name string `json:"name"`
	PosX uint   `json:"pos_x"`
	PosY uint   `json:"pos_y"`

	AccountID uint    `json:"account_id"`
	Account   Account `json:"account"`

	TicketID uint   `json:"ticket_id"`
	Ticket   Ticket `json:"ticket"`
}
