package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type TicketGormRepository struct {
	database *gorm.DB
}

func (r *TicketGormRepository) List() ([]domain.Ticket, *domain.AppError) {
	var tickets []domain.Ticket
	r.database.
		Preload("Account").
		Preload("TicketProducts").
		Find(&tickets)
	return tickets, nil
}

func (r *TicketGormRepository) Save(user *domain.Ticket) (*domain.Ticket, *domain.AppError) {
	result := r.database.Save(user)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear usuario")
	}
	return user, nil
}

func (r *TicketGormRepository) Find(id uint) (*domain.Ticket, *domain.AppError) {
	ticket := new(domain.Ticket)
	r.database.
		Preload("Account").
		Preload("TicketProducts").
		First(ticket, id)
	if ticket.ID == 0 {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
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
	r.database.Delete(ticket)
	return ticket, nil
}

func NewTicketGormRepository(database *gorm.DB) *TicketGormRepository {
	database.AutoMigrate(&domain.Ticket{})
	return &TicketGormRepository{database}
}
