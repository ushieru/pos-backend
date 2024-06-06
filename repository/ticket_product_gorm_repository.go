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

func (r *TicketProductGormRepository) List(criteria *domain_criteria.Criteria) ([]domain.TicketProduct, *domain.AppError) {
	var ticketProducts []domain.TicketProduct
	scopes := r.c.FiltersToScopes(criteria.Filters)
	tx := r.database.Model(&domain.TicketProduct{})
	if len(scopes) > 0 {
		tx.Scopes(scopes...)
	}
	tx.Find(&ticketProducts)
	return ticketProducts, nil
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

func (r *TicketProductGormRepository) FindByAccountProductionCenters(account *domain.Account) ([]domain.TicketProduct, *domain.AppError) {
	var ticketProducts []domain.TicketProduct
	r.database.
		Model(&domain.TicketProduct{}).
		Select("ticket_products.*").
		Joins("INNER JOIN category_product ON ticket_products.product_id = category_product.product_id").
		Joins("INNER JOIN categories ON category_product.category_id = categories.id").
		Joins("INNER JOIN account_production_center ON account_production_center.account_id = ?", account.ID).
		Where("categories.production_center_id = account_production_center.production_center_id").
		Where("ticket_products.status <> ? AND ticket_products.status <> ?", domain.Added, domain.Prepared).
		Scan(&ticketProducts)
	return ticketProducts, nil
}

func (r *TicketProductGormRepository) Update(ticketProduct *domain.TicketProduct) (*domain.TicketProduct, *domain.AppError) {
	r.database.Save(ticketProduct)
	return ticketProduct, nil
}

func NewTicketProductGormRepository(database *gorm.DB) domain.ITicketProductRepository {
	database.AutoMigrate(&domain.TicketProduct{})
	return &TicketProductGormRepository{database: database}
}
