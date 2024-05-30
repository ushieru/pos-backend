package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type TableGormRepository struct {
	database *gorm.DB
}

func (r *TableGormRepository) List() ([]domain.Table, *domain.AppError) {
	var table []domain.Table
	r.database.
		Preload("Account").
		Preload("Ticket").
		Find(&table)
	return table, nil
}

func (r *TableGormRepository) Save(table *domain.Table) (*domain.Table, *domain.AppError) {
	tables, err := r.List()
	if err != nil {
		return nil, err
	}
	var tables2dSlice = make([][]domain.Table, 5)
	for i := range tables2dSlice {
		tables2dSlice[i] = make([]domain.Table, 10)
	}
	for _, _table := range tables {
		tables2dSlice[table.PosY][table.PosX] = _table
	}
	var posXAvailable, posYAvailable uint
	findPos := false
	for y, rt := range tables2dSlice {
		for x, t := range rt {
			if t.ID == "" {
				posXAvailable, posYAvailable = uint(x+1), uint(y+1)
				findPos = true
				break
			}
		}
		if findPos {
			break
		}
	}
	table.PosX = posXAvailable
	table.PosY = posYAvailable
	result := r.database.Save(table)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear Mesa")
	}
	return table, nil
}

func (r *TableGormRepository) Find(id string) (*domain.Table, *domain.AppError) {
	table := new(domain.Table)
	r.database.
		Preload("Account").
		Preload("Ticket").
		First(table, id)
	if table.ID == "" {
		return nil, domain.NewNotFoundError("Mesa no encontrada")
	}
	return table, nil
}

func (r *TableGormRepository) Update(t *domain.Table) (*domain.Table, *domain.AppError) {
	table, err := r.Find(t.ID)
	if err != nil {
		return nil, err
	}
	tableFind := new(domain.Table)
	r.database.First(tableFind, "Pos_X = ? AND Pos_Y = ?", t.PosX, t.PosY)
	if tableFind.ID != "0" && tableFind.ID != t.ID {
		return nil, domain.NewConflictError("Posicion ocupada")
	}
	table.Name = t.Name
	table.PosX = t.PosX
	table.PosY = t.PosY
	r.database.Save(table)
	return table, nil
}

func (r *TableGormRepository) Delete(id string) (*domain.Table, *domain.AppError) {
	table, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(table)
	return table, nil
}

func (r *TableGormRepository) CreateTicket(
	t *domain.Table,
	a *domain.Account,
) (*domain.Table, *domain.AppError) {
	table, err := r.Find(t.ID)
	if err != nil {
		return nil, err
	}
	ticket := &domain.Ticket{
		Account:      *a,
		TicketStatus: domain.TicketOpen,
	}
	r.database.Save(ticket)
	table.AccountID = a.ID
	table.TicketID = ticket.ID
	r.database.Save(table)
	return table, nil
}

func NewTableGormRepository(database *gorm.DB) domain.ITableRepository {
	database.AutoMigrate(&domain.Table{})
	return &TableGormRepository{database}
}
