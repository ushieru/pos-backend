package service

import (
	"github.com/ushieru/pos/domain"
)

type ITicketService interface {
	List() ([]domain.Ticket, *domain.AppError)
	Find(id uint) (*domain.Ticket, *domain.AppError)
	Save(account *domain.Account) (*domain.Ticket, *domain.AppError)
	Delete(id uint) (*domain.Ticket, *domain.AppError)
}

type TicketService struct {
	repository domain.ITicketRepository
}

func (s *TicketService) List() ([]domain.Ticket, *domain.AppError) {
	return s.repository.List()
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

func NewTicketService(repository domain.ITicketRepository) *TicketService {
	return &TicketService{repository}
}