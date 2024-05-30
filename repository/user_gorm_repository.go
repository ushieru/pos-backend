package repository

import (
	"github.com/ushieru/pos/domain"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	database *gorm.DB
}

func (r *UserGormRepository) List() ([]domain.User, *domain.AppError) {
	var user []domain.User
	r.database.Preload("Account").Find(&user)
	return user, nil
}

func (r *UserGormRepository) Save(user *domain.User) (*domain.User, *domain.AppError) {
	result := r.database.Save(user)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear usuario")
	}
	return user, nil
}

func (r *UserGormRepository) Find(id string) (*domain.User, *domain.AppError) {
	user := new(domain.User)
	r.database.Preload("Account").First(user, "id = ?", id)
	if user.ID == "" {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) FindByUserOrEmailAndPassword(username, password string) (*domain.User, *domain.AppError) {
	user := new(domain.User)
	r.database.Joins("Account").First(user, "(Email = ? or Account.Username = ?) AND password = ?", username, username, password)
	if user.ID == "" {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) Update(user *domain.User) (*domain.User, *domain.AppError) {
	r.database.Save(user)
	r.database.Model(&domain.Account{}).Save(user.Account)
	return user, nil
}

func (r *UserGormRepository) Delete(user *domain.User) (*domain.User, *domain.AppError) {
	r.database.Delete(user)
	return user, nil
}

func NewUserGormRepository(database *gorm.DB) domain.IUserRepository {
	database.AutoMigrate(&domain.User{})
	r := UserGormRepository{database}
	return &r
}
