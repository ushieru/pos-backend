package repository

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/domain/criteria"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TicketGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
}

func (r *TicketGormRepository) List(c *domain_criteria.Criteria) ([]domain.Ticket, *domain.AppError) {
	var tickets []domain.Ticket
	scopes := r.c.FiltersToScopes(c.Filters)
	statement := r.database.Model(&domain.Ticket{})
	if len(scopes) > 0 {
		statement = statement.Scopes(scopes...)
	}
	statement.
		Preload(clause.Associations).
		Find(&tickets)
	if tickets == nil {
		tickets = make([]domain.Ticket, 0)
	}
	return tickets, nil
}

func (r *TicketGormRepository) Save(ticket *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	result := r.database.Save(ticket)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear ticket")
	}
	return ticket, nil
}

func (r *TicketGormRepository) Update(ticket *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	result := r.database.Save(ticket)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al actualizar ticket")
	}
	return ticket, nil
}

func (r *TicketGormRepository) OrderTicketProductsByTicket(ticket *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	r.database.
		Model(&domain.TicketProduct{}).
		Where("ticket_id = ? AND status = ?", ticket.ID, "Added").
		Update("status", domain.Ordered)
	updatedTicket, _ := r.Find(ticket.ID)
	return updatedTicket, nil
}

func (r *TicketGormRepository) Find(id string) (*domain.Ticket, *domain.AppError) {
	ticket := new(domain.Ticket)
	r.database.
		Preload("Account").
		Preload("TicketProducts").
		First(ticket, "id = ?", id)
	if ticket.ID == "" {
		return nil, domain.NewNotFoundError("Ticket no encontrado")
	}
	return ticket, nil
}

func (r *TicketGormRepository) Delete(ticket *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	r.database.Delete(ticket)
	return ticket, nil
}

func (r *TicketGormRepository) AddProduct(ticket *domain.Ticket, product *domain.Product) (*domain.Ticket, *domain.AppError) {
	ticketProduct := &domain.TicketProduct{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Status:      domain.Added,
		ProductID:   product.ID,
		TicketID:    ticket.ID,
	}
	r.database.
		Model(ticketProduct).
		Create(ticketProduct)
	total := product.Price
	for _, productTicket := range ticket.TicketProducts {
		total += productTicket.Price
	}
	ticket.Total = total
	r.database.Save(ticket)
	return ticket, nil
}

func (r *TicketGormRepository) DeleteProduct(ticket *domain.Ticket, product *domain.Product) (*domain.Ticket, *domain.AppError) {
	ticketProduct := new(domain.TicketProduct)
	r.database.
		Model(ticketProduct).
		Where("product_id = ? AND status = ?", product.ID, domain.Added).
		First(ticketProduct)
	if ticketProduct.ID == "" {
		return nil, domain.NewUnauthorizedError("Producto en ticket no encontrado")
	}
	r.database.Delete(ticketProduct)
	updatedTicket, UpdatedTicketErr := r.Find(ticket.ID)
	if UpdatedTicketErr != nil {
		return nil, UpdatedTicketErr
	}
	total := 0.0
	for _, productTicket := range updatedTicket.TicketProducts {
		total += productTicket.Price
	}
	updatedTicket.Total = total
	r.database.Save(updatedTicket)
	return updatedTicket, nil
}

func (r *TicketGormRepository) PayTicket(ticket *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	ticket.TicketStatus = domain.TicketPaid
	r.database.Save(ticket)
	return ticket, nil
}

func NewTicketGormRepository(database *gorm.DB) domain.ITicketRepository {
	database.AutoMigrate(&domain.Ticket{})
	database.AutoMigrate(&domain.TicketProduct{})
	return &TicketGormRepository{database: database}
}
