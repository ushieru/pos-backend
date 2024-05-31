package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type ITicketService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.Ticket, *domain.AppError)
	Find(id string) (*domain.Ticket, *domain.AppError)
	Save(account *domain.Account) (*domain.Ticket, *domain.AppError)
	Delete(id string) (*domain.Ticket, *domain.AppError)
	AddProduct(ticketId, productId string, a *domain.Account) (*domain.Ticket, *domain.AppError)
	DeleteProduct(ticketId, productId string, a *domain.Account) (*domain.Ticket, *domain.AppError)
	ChangeTable(ticketId, tableId string) (*domain.Ticket, *domain.AppError)
	PayTicket(id string, a *domain.Account) (*domain.Ticket, *domain.AppError)
	OrderTicketProducts(id string, a *domain.Account) (*domain.Ticket, *domain.AppError)
}

type TicketService struct {
	ticketRepository  domain.ITicketRepository
	tableRepository   domain.ITableRepository
	productRepository domain.IProductRepository
}

func (s *TicketService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.Ticket, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.ticketRepository.List(criteria)
}

func (s *TicketService) Find(id string) (*domain.Ticket, *domain.AppError) {
	return s.ticketRepository.Find(id)
}

func (s *TicketService) Save(account *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket := &domain.Ticket{
		TicketStatus: domain.TicketOpen,
		Account:      *account,
	}
	return s.ticketRepository.Save(ticket)
}

func (s *TicketService) Delete(id string) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	if len(ticket.TicketProducts) != 0 {
		return nil, domain.NewConflictError("Ticket no esta vacio")
	}
	table, _ := s.tableRepository.FindByTicketId(id)
	if table != nil {
		if _, err := s.tableRepository.UpdateAccountRelation(table, nil); err != nil {
			return nil, err
		}
		if _, err := s.tableRepository.UpdateTicketRelation(table, nil); err != nil {
			return nil, err
		}
	}
	return s.ticketRepository.Delete(ticket)
}

func (s *TicketService) AddProduct(ticketId, productId string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.Find(ticketId)
	if err != nil {
		return nil, err
	}
	if ticket.AccountID != a.ID {
		return nil, domain.NewUnauthorizedError("Este ticket no es tuyo")
	}
	if ticket.TicketStatus != domain.TicketOpen {
		return nil, domain.NewConflictError("Este ticket no esta abierto")
	}
	product, err := s.productRepository.
		Find(productId, &domain_criteria.Criteria{
			Filters: []domain_criteria.Filter{
				{
					Field:    "instr(available_days, strftime('%w', date('now')))",
					Operator: domain_criteria.GT,
					Value:    "0",
				},
				{
					Field:    "time('now', 'localtime')",
					Operator: domain_criteria.GTE,
					Value:    "time(available_from_hour)",
				},
				{
					Field:    "time('now', 'localtime')",
					Operator: domain_criteria.LTE,
					Value:    "time(available_until_hour)",
				},
			},
		})
	if err != nil {
		return nil, err
	}
	return s.ticketRepository.AddProduct(ticket, product)
}

func (s *TicketService) DeleteProduct(ticketId, productId string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.Find(ticketId)
	if err != nil {
		return nil, err
	}
	if a.AccountType != domain.Admin {
		if ticket.AccountID != a.ID {
			return nil, domain.NewUnauthorizedError("Este ticket no es tuyo")
		}
	}
	if ticket.TicketStatus != domain.TicketOpen {
		return nil, domain.NewConflictError("Este ticket no esta abierto")
	}
	product, err := s.productRepository.Find(productId, nil)
	if err != nil {
		return nil, err
	}
	return s.ticketRepository.DeleteProduct(ticket, product)
}

func (s *TicketService) PayTicket(ticketId string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	if a.AccountType != domain.Cashier {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para cobrar")
	}
	ticket, err := s.Find(ticketId)
	if err != nil {
		return nil, err
	}
	if len(ticket.TicketProducts) == 0 {
		return nil, domain.NewConflictError("Ticket vacio")
	}
	table, _ := s.tableRepository.FindByTicketId(ticketId)
	if table != nil {
		if _, err := s.tableRepository.UpdateAccountRelation(table, nil); err != nil {
			return nil, err
		}
		if _, err := s.tableRepository.UpdateTicketRelation(table, nil); err != nil {
			return nil, err
		}
	}
	s.ticketRepository.OrderTicketProductsByTicket(ticket)
	return s.ticketRepository.PayTicket(ticket)
}

func (s *TicketService) ChangeTable(ticketId, tableId string) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.ticketRepository.Find(ticketId)
	if err != nil {
		return nil, err
	}
	table, err := s.tableRepository.FindByTicketId(ticketId)
	if err != nil {
		return nil, err
	}
	if _, err := s.tableRepository.UpdateTicketRelation(table, ticket); err != nil {
		return nil, err
	}
	updatedTicket, _ := s.ticketRepository.Find(ticketId)
	return updatedTicket, nil
}

func (s *TicketService) OrderTicketProducts(id string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.ticketRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return s.ticketRepository.OrderTicketProductsByTicket(ticket)
}

func NewTicketService(
	ticketRepository domain.ITicketRepository,
	tableRepository domain.ITableRepository,
	productRepository domain.IProductRepository,
) *TicketService {
	return &TicketService{
		ticketRepository,
		tableRepository,
		productRepository,
	}
}
