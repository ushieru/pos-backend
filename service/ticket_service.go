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
	table, err := s.tableRepository.FindByTicketId(id)
	if err != nil {
		return nil, err
	}
	if _, err := s.tableRepository.UpdateAccountRelation(table, nil); err != nil {
		return nil, err
	}
	if _, err := s.tableRepository.UpdateTicketRelation(table, nil); err != nil {
		return nil, err
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
	product, err := s.productRepository.Find(productId)
	if err != nil {
		return nil, err
	}
	// TODO: Add validation in availability dates
	return s.ticketRepository.AddProduct(ticket, product)
}

func (s *TicketService) DeleteProduct(ticketId, productId string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
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
	product, pError := s.productRepository.Find(productId)
	if pError != nil {
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
	table, err := s.tableRepository.FindByTicketId(ticketId)
	if err != nil {
		return nil, err
	}
	if _, err := s.tableRepository.UpdateAccountRelation(table, nil); err != nil {
		return nil, err
	}
	if _, err := s.tableRepository.UpdateTicketRelation(table, nil); err != nil {
		return nil, err
	}
	return s.ticketRepository.PayTicket(ticket)
}

func (s *TicketService) OrderTicketProducts(id string, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket, err := s.ticketRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return s.ticketRepository.OrderTicketProductsByTicket(ticket)
}

func NewTicketService(ticketRepository domain.ITicketRepository, tableRepository domain.ITableRepository) *TicketService {
	return &TicketService{
		ticketRepository: ticketRepository,
		tableRepository:  tableRepository,
	}
}
