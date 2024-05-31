package service

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
)

type ITableService interface {
	List() ([]domain.Table, *domain.AppError)
	Find(id string) (*domain.Table, *domain.AppError)
	Save(dto *dto.CreateTableRequest) (*domain.Table, *domain.AppError)
	Update(id string, dto *dto.UpdateTableRequest) (*domain.Table, *domain.AppError)
	Delete(id string) (*domain.Table, *domain.AppError)
	CreateTicket(tableId string, account *domain.Account) (*domain.Table, *domain.AppError)
}

type TableService struct {
	repository domain.ITableRepository
}

func (s *TableService) List() ([]domain.Table, *domain.AppError) {
	return s.repository.List()
}

func (s *TableService) Find(id string) (*domain.Table, *domain.AppError) {
	return s.repository.Find(id)
}

func (s *TableService) Save(dto *dto.CreateTableRequest) (*domain.Table, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	table := &domain.Table{Name: dto.Name}
	return s.repository.Save(table)
}

func (s *TableService) Update(id string, dto *dto.UpdateTableRequest) (*domain.Table, *domain.AppError) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	table := &domain.Table{
		Model: domain.Model{ID: id},
		Name:  dto.Name,
		PosX:  dto.PosX,
		PosY:  dto.PosY,
	}
	return s.repository.Update(table)
}

func (s *TableService) Delete(id string) (*domain.Table, *domain.AppError) {
	table, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	return s.repository.Delete(table)
}
func (s *TableService) CreateTicket(tableId string, account *domain.Account) (*domain.Table, *domain.AppError) {
	table := &domain.Table{
		Model: domain.Model{ID: tableId},
	}
	return s.repository.CreateTicket(table, account)
}

func NewTableService(repository domain.ITableRepository) *TableService {
	return &TableService{repository}
}
