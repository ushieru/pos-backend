package domain

type TicketStatus string

const (
	TicketOpen   TicketStatus = "open"
	TicketCancel              = "cancel"
	TicketPaid                = "paid"
)

type Ticket struct {
	Model
	TicketStatus   TicketStatus    `json:"ticket_status"`
	Total          float64         `json:"total" gorm:"default:0"`
	AccountID      uint            `json:"account_id"`
	Account        Account         `json:"account"`
	TicketProducts []TicketProduct `json:"ticket_products"`
}

type ITicketRepository interface {
	List() ([]Ticket, *AppError)
	Save(*Ticket) (*Ticket, *AppError)
	Find(id uint) (*Ticket, *AppError)
	Delete(id uint) (*Ticket, *AppError)
}
