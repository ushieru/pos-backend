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
	return nil, domain.NewUnexpectedError("[TableGormRepository] (Save) - Metodo no implementado")
}

func (r *TableGormRepository) Find(id uint) (*domain.Table, *domain.AppError) {
	table := new(domain.Table)
	r.database.
		Preload("Account").
		Preload("Ticket").
		First(table, id)
	if table.ID == 0 {
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
	r.database.First(tableFind, "pos_x = ? AND pos_y = ?", t.PosX, t.PosY)
	if tableFind.ID != 0 {
		return nil, domain.NewConflictError("Posicion ocupada")
	}
	table.Name = t.Name
	table.PosX = t.PosX
	table.PosY = t.PosY
	r.database.Save(table)
	return table, nil
}

func (r *TableGormRepository) Delete(id uint) (*domain.Table, *domain.AppError) {
	table, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(table)
	return table, nil
}

func NewTableGormRepository(database *gorm.DB) *TableGormRepository {
	database.AutoMigrate(&domain.Table{})
	return &TableGormRepository{database}
}
