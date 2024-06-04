package service

import (
	"github.com/ushieru/pos/domain"
	domain_criteria "github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/dto"
)

type IProductionCenterService interface {
	List(*dto.SearchCriteriaQueryRequest) ([]domain.ProductionCenter, *domain.AppError)
	Find(id string) (*domain.ProductionCenter, *domain.AppError)
	Save(dto *dto.CreateProductionCenterRequest) (*domain.ProductionCenter, *domain.AppError)
	Update(id string, dto *dto.CreateProductionCenterRequest) (*domain.ProductionCenter, *domain.AppError)
	GetTicket(id, ticketId string) (*domain.Ticket, *domain.AppError)
	AddAccount(productionCenterId, accountId string) (*domain.ProductionCenter, *domain.AppError)
	DeleteAccount(productionCenterId, accountId string) (*domain.ProductionCenter, *domain.AppError)
	AddCategory(productionCenterId, categoryId string) (*domain.ProductionCenter, *domain.AppError)
	DeleteCategory(productionCenterId, categoryId string) (*domain.ProductionCenter, *domain.AppError)
	Delete(id string) (*domain.ProductionCenter, *domain.AppError)
}

type ProductionCenterService struct {
	productionCenterRepository domain.IProductionCenterRepository
	accountRepository          domain.IAccountRepository
	categoryRepository         domain.ICategoryRepository
	ticketRepository           domain.ITicketRepository
	ticketProductRepository    domain.ITicketProductRepository
}

func (c *ProductionCenterService) List(dto *dto.SearchCriteriaQueryRequest) ([]domain.ProductionCenter, *domain.AppError) {
	criteria := &domain_criteria.Criteria{
		Filters: dto.Filters,
	}
	return c.productionCenterRepository.List(criteria)
}

func (c *ProductionCenterService) Find(id string) (*domain.ProductionCenter, *domain.AppError) {
	return c.productionCenterRepository.Find(id)
}

func (c *ProductionCenterService) Save(dto *dto.CreateProductionCenterRequest) (*domain.ProductionCenter, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	productionCenter := &domain.ProductionCenter{
		Name: dto.Name,
	}
	return c.productionCenterRepository.Save(productionCenter)

}

func (c *ProductionCenterService) Update(id string, dto *dto.CreateProductionCenterRequest) (*domain.ProductionCenter, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	productionCenter, err := c.productionCenterRepository.Find(id)
	if err != nil {
		return nil, err
	}
	productionCenter.Name = dto.Name
	return c.productionCenterRepository.Update(productionCenter)
}

func (c *ProductionCenterService) GetTicket(id, ticketId string) (*domain.Ticket, *domain.AppError) {
	productionCenter, err := c.productionCenterRepository.Find(id)
	if err != nil {
		return nil, err
	}
	ticket, err := c.ticketRepository.Find(ticketId)
	if err != nil {
		return nil, err
	}
	ticketProducst, err := c.ticketProductRepository.FindByProductionCenter(productionCenter.ID)
	if err != nil {
		return nil, err
	}
	ticket.TicketProducts = ticketProducst
	return ticket, nil
}

func (c *ProductionCenterService) AddAccount(productionCenterId, accountId string) (*domain.ProductionCenter, *domain.AppError) {
	productionCenter, err := c.productionCenterRepository.Find(productionCenterId)
	if err != nil {
		return nil, err
	}
	account, err := c.accountRepository.Find(accountId)
	if err != nil {
		return nil, err
	}
	return c.productionCenterRepository.AddAccount(productionCenter, account)
}

func (c *ProductionCenterService) DeleteAccount(productionCenterId, accountId string) (*domain.ProductionCenter, *domain.AppError) {
	productionCenter, err := c.productionCenterRepository.Find(productionCenterId)
	if err != nil {
		return nil, err
	}
	account, err := c.accountRepository.Find(accountId)
	if err != nil {
		return nil, err
	}
	return c.productionCenterRepository.DeleteAccount(productionCenter, account)
}

func (c *ProductionCenterService) AddCategory(productionCenterId, categoryId string) (*domain.ProductionCenter, *domain.AppError) {
	productionCenter, err := c.productionCenterRepository.Find(productionCenterId)
	if err != nil {
		return nil, err
	}
	category, err := c.categoryRepository.Find(categoryId)
	if err != nil {
		return nil, err
	}
	return c.productionCenterRepository.AddCategory(productionCenter, category)
}

func (c *ProductionCenterService) DeleteCategory(productionCenterId, categoryId string) (*domain.ProductionCenter, *domain.AppError) {
	productionCenter, err := c.productionCenterRepository.Find(productionCenterId)
	if err != nil {
		return nil, err
	}
	category, err := c.categoryRepository.Find(categoryId)
	if err != nil {
		return nil, err
	}
	return c.productionCenterRepository.DeleteCategory(productionCenter, category)
}

func (c *ProductionCenterService) Delete(id string) (*domain.ProductionCenter, *domain.AppError) {
	category, err := c.productionCenterRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return c.productionCenterRepository.Delete(category)
}

func NewProductionCenterService(
	productionCenterRepository domain.IProductionCenterRepository,
	accountRepository domain.IAccountRepository,
	categoryRepository domain.ICategoryRepository,
	ticketRepository domain.ITicketRepository,
	ticketProductRepository domain.ITicketProductRepository,
) IProductionCenterService {
	return &ProductionCenterService{
		productionCenterRepository,
		accountRepository,
		categoryRepository,
		ticketRepository,
		ticketProductRepository,
	}
}
