package repository

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
)

type TicketProductGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *TicketProductGormRepository) Find(id string, criteria *domain_criteria.Criteria) (*domain.TicketProduct, *domain.AppError) {
	ticketProduct := new(domain.TicketProduct)
	stm := r.database.Model(ticketProduct)
	if criteria == nil {
		criteria = &domain_criteria.Criteria{}
	}
	scopes := r.c.FiltersToScopes(criteria.Filters)
	if len(scopes) > 0 {
		stm.Scopes(scopes...)
	}
	stm.First(ticketProduct, "id = ?", id)
	if ticketProduct.ID == "" {
		return nil, domain.NewNotFoundError("Ticket Product no encontrado")
	}
	return ticketProduct, nil
}

func (r *TicketProductGormRepository) Update(ticketProduct *domain.TicketProduct) (*domain.TicketProduct, *domain.AppError) {
	r.database.Save(ticketProduct)
	return ticketProduct, nil
}

func NewTicketProductGormRepository(database *gorm.DB) domain.ITicketProductRepository {
	database.AutoMigrate(&domain.TicketProduct{})
	return &TicketProductGormRepository{database: database}
}
