package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type AccountGormRepository struct {
	database *gorm.DB
}

func (r *AccountGormRepository) Find(id string) (*domain.Account, *domain.AppError) {
	account := new(domain.Account)
	r.database.First(account, "id = ?", id)
	if account.ID == "" {
		return nil, domain.NewNotFoundError("Cuenta no encontrada")
	}
	return account, nil
}

func NewAccountGormRepository(database *gorm.DB) domain.IAccountRepository {
	database.AutoMigrate(&domain.Account{})
	r := AccountGormRepository{database: database}
	return &r
}
