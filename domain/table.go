package domain

type Table struct {
	Model
	Name      string  `json:"name"`
	PosX      uint    `json:"pos_x"`
	PosY      uint    `json:"pos_y"`
	AccountID string  `json:"account_id"`
	Account   Account `json:"account"`
	TicketID  string  `json:"ticket_id"`
	Ticket    Ticket  `json:"ticket"`
}

type ITableRepository interface {
	List() ([]Table, *AppError)
	Save(*Table) (*Table, *AppError)
	Find(id string) (*Table, *AppError)
	FindByTicketId(id string) (*Table, *AppError)
	UpdateAccountRelation(*Table, *Account) (*Table, *AppError)
	UpdateTicketRelation(*Table, *Ticket) (*Table, *AppError)
	Update(*Table) (*Table, *AppError)
	Delete(*Table) (*Table, *AppError)
	CreateTicket(*Table, *Account) (*Table, *AppError)
}
