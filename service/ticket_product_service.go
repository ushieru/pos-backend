package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type ITicketProductService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.TicketProduct, *domain.AppError)
	FindByAccountProductionCenters(*domain.Account) ([]domain.TicketProduct, *domain.AppError)
	InPreparation(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError)
	Prepared(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError)
	Paid(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError)
}

type TicketProductService struct {
	repository                 domain.ITicketProductRepository
	productionCenterRepository domain.IProductionCenterRepository
}

func (s *TicketProductService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.TicketProduct, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return s.repository.List(criteria)
}

func (s *TicketProductService) FindByAccountProductionCenters(account *domain.Account) ([]domain.TicketProduct, *domain.AppError) {
	return s.repository.FindByAccountProductionCenters(account)
}

func (s *TicketProductService) InPreparation(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError) {
	if account.AccountType != domain.Producer {
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
	if account.AccountType != domain.Producer {
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

func (s *TicketProductService) Paid(ticketProductId string, account *domain.Account) (*domain.TicketProduct, *domain.AppError) {
	if account.AccountType != domain.Cashier {
		return nil, domain.NewUnauthorizedError("No estas autorizado para esta accion")
	}
	ticketProduct, err := s.repository.Find(ticketProductId, &domain_criteria.Criteria{})
	if err != nil {
		return nil, err
	}
	ticketProduct.Status = domain.Paid
	// TODO: recalculate total ticket
	return s.repository.Update(ticketProduct)
}

func NewTicketProductService(repository domain.ITicketProductRepository, productionCenterRepository domain.IProductionCenterRepository) *TicketProductService {
	return &TicketProductService{repository, productionCenterRepository}
}
