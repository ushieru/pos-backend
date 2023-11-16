package repository

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/domain/criteria"
	"github.com/ushieru/pos/service"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TicketGormRepository struct {
	c        CriteriaGormRepository
	database *gorm.DB
	ps       service.IProductService
}

func (r *TicketGormRepository) List(
	c *domain_criteria.Criteria,
) ([]domain.Ticket, *domain.AppError) {
	var tickets []domain.Ticket
	scopes := r.c.FiltersToScopes(c.Filters)
	statement := r.database.Debug()
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

func (r *TicketGormRepository) Find(id uint) (*domain.Ticket, *domain.AppError) {
	ticket := new(domain.Ticket)
	r.database.
		Preload("Account").
		Preload("TicketProducts").
		First(ticket, id)
	if ticket.ID == 0 {
		return nil, domain.NewNotFoundError("Ticket no encontrado")
	}
	return ticket, nil
}

func (r *TicketGormRepository) Delete(id uint) (*domain.Ticket, *domain.AppError) {
	ticket, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	if len(ticket.TicketProducts) != 0 {
		return nil, domain.NewConflictError("Ticket no esta vacio")
	}
	table := new(domain.Table)
	r.database.First(table, "ticket_id = ?", id)
	if table.ID != 0 {
		r.database.Model(&table).Association("Account").Clear()
		r.database.Model(&table).Association("Ticket").Clear()
	}
	r.database.Delete(ticket)
	return ticket, nil
}

func (r *TicketGormRepository) AddProduct(
	ticketId, productId uint,
	a *domain.Account,
) (*domain.Ticket, *domain.AppError) {
	ticket, err := r.Find(ticketId)
	if err != nil {
		return nil, err
	}
	if ticket.AccountID != a.ID {
		return nil, domain.NewUnauthorizedError("Este ticket no es tuyo")
	}
	if ticket.TicketStatus != domain.TicketOpen {
		return nil, domain.NewConflictError("Este ticket no esta abierto")
	}
	product, pError := r.ps.Find(productId)
	if pError != nil {
		return nil, err
	}
	// TODO: Add validation in availability dates
	tp := new(domain.TicketProduct)
	r.database.First(tp, "Product_Id = ? AND ticket_id = ?", productId, ticketId)
	if tp.ID == 0 {
		ticketProduct := new(domain.TicketProduct)
		ticketProduct.Quantity = 1
		ticketProduct.Name = product.Name
		ticketProduct.Description = product.Description
		ticketProduct.Price = product.Price
		ticketProduct.ProductId = product.ID
		r.database.Model(ticket).Association("TicketProducts").Append(ticketProduct)
	}
	if tp.ID != 0 {
		tp.Quantity = tp.Quantity + 1
		r.database.Save(tp)
	}
	updatedTicket, UpdatedTicketErr := r.Find(ticketId)
	if UpdatedTicketErr != nil {
		return nil, UpdatedTicketErr
	}
	total := 0.0
	for _, productTicket := range updatedTicket.TicketProducts {
		total += productTicket.Product.Price * float64(productTicket.Quantity)
	}
	updatedTicket.Total = total
	r.database.Save(updatedTicket)
	return updatedTicket, nil
}

func (r *TicketGormRepository) DeleteProduct(
	ticketId, productId uint,
	a *domain.Account,
) (*domain.Ticket, *domain.AppError) {
	ticket, err := r.Find(ticketId)
	if err != nil {
		return nil, err
	}
	if ticket.AccountID != a.ID {
		return nil, domain.NewUnauthorizedError("Este ticket no es tuyo")
	}
	if ticket.TicketStatus != domain.TicketOpen {
		return nil, domain.NewConflictError("Este ticket no esta abierto")
	}
	product, pError := r.ps.Find(productId)
	if pError != nil || product.ID == 0 {
		return nil, err
	}
	tp := new(domain.TicketProduct)
	r.database.First(tp, "Product_Id = ? AND ticket_id = ?", productId, ticketId)
	if tp.Quantity == 1 {
		r.database.Model(ticket).Association("TicketProducts").Delete(tp)
		r.database.Unscoped().Delete(tp)
	}
	if tp.Quantity > 1 {
		tp.Quantity = tp.Quantity - 1
		r.database.Save(tp)
	}
	updatedTicket, UpdatedTicketErr := r.Find(ticketId)
	if UpdatedTicketErr != nil {
		return nil, UpdatedTicketErr
	}
	total := 0.0
	for _, productTicket := range updatedTicket.TicketProducts {
		total += productTicket.Product.Price * float64(productTicket.Quantity)
	}
	updatedTicket.Total = total
	r.database.Save(updatedTicket)
	return updatedTicket, nil
}

func (r *TicketGormRepository) PayTicket(
	id uint,
	a *domain.Account,
) (*domain.Ticket, *domain.AppError) {
	if a.AccountType != domain.Cashier {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para cobrar")
	}
	ticket, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	if len(ticket.TicketProducts) == 0 {
		return nil, domain.NewConflictError("Ticket vacio")
	}
	table := new(domain.Table)
	r.database.First(table, "ticket_id = ?", id)
	if table.ID != 0 {
		r.database.Model(table).Association("Account").Clear()
		r.database.Model(table).Association("Ticket").Clear()
	}
	ticket.TicketStatus = domain.TicketPaid
	r.database.Save(ticket)
	return ticket, nil
}

func NewTicketGormRepository(
	database *gorm.DB,
	ps service.IProductService,
) domain.ITicketRepository {
	database.AutoMigrate(&domain.Ticket{})
	database.AutoMigrate(&domain.TicketProduct{})
	return &TicketGormRepository{database: database, ps: ps}
}
