package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type ITicketService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.Ticket, *domain.AppError)
	Find(id uint) (*domain.Ticket, *domain.AppError)
	Save(account *domain.Account) (*domain.Ticket, *domain.AppError)
	Delete(id uint) (*domain.Ticket, *domain.AppError)
	AddProduct(ticketId, productId uint, a *domain.Account) (*domain.Ticket, *domain.AppError)
	DeleteProduct(ticketId, productId uint, a *domain.Account) (*domain.Ticket, *domain.AppError)
	PayTicket(id uint, a *domain.Account) (*domain.Ticket, *domain.AppError)
}

type TicketService struct {
	repository domain.ITicketRepository
}

func (s *TicketService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.Ticket, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.repository.List(criteria)
}

func (s *TicketService) Find(id uint) (*domain.Ticket, *domain.AppError) {
	return s.repository.Find(id)
}

func (s *TicketService) Save(account *domain.Account) (*domain.Ticket, *domain.AppError) {
	ticket := &domain.Ticket{
		TicketStatus: domain.TicketOpen,
		Account:      *account,
	}
	return s.repository.Save(ticket)
}

func (s *TicketService) Delete(id uint) (*domain.Ticket, *domain.AppError) {
	return s.repository.Delete(id)
}

func (s *TicketService) AddProduct(ticketId, productId uint, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	return s.repository.AddProduct(ticketId, productId, a)
}

func (s *TicketService) DeleteProduct(ticketId, productId uint, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	return s.repository.DeleteProduct(ticketId, productId, a)
}
func (s *TicketService) PayTicket(id uint, a *domain.Account) (*domain.Ticket, *domain.AppError) {
	return s.repository.PayTicket(id, a)
}

func NewTicketService(repository domain.ITicketRepository) *TicketService {
	return &TicketService{repository}
}
