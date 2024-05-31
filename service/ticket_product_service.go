package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
)

type ITicketProductService interface {
	InPreparation(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError)
	Prepared(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError)
}

type TicketProductService struct {
	repository domain.ITicketProductRepository
}

func (s *TicketProductService) InPreparation(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError) {
	if account.AccountType != domain.Cook && account.AccountType != domain.Bartender {
		return nil, domain.NewUnauthorizedError("No estas autorizado para esta accion")
	}
	ticketProduct, err := s.repository.Find(ticketProductId, &domain_criteria.Criteria{
		Filters: []domain_criteria.Filter{
			{
				Field:    "status",
				Operator: domain_criteria.NOT_EQUAL,
				Value:    "'Added'",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	ticketProduct.Status = domain.InPreparation
	return s.repository.Update(ticketProduct)
}

func (s *TicketProductService) Prepared(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError) {
	if account.AccountType != domain.Cook && account.AccountType != domain.Bartender {
		return nil, domain.NewUnauthorizedError("No estas autorizado para esta accion")
	}
	ticketProduct, err := s.repository.Find(ticketProductId, &domain_criteria.Criteria{
		Filters: []domain_criteria.Filter{
			{
				Field:    "status",
				Operator: domain_criteria.NOT_EQUAL,
				Value:    "'Added'",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	ticketProduct.Status = domain.Prepared
	return s.repository.Update(ticketProduct)
}

func NewTicketProductService(repository domain.ITicketProductRepository) *TicketProductService {
	return &TicketProductService{repository}
}
