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
	r.database.Preload("Account").First(user, id)
	if user.ID == "" {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) FindByUserOrEmail(username string) (*domain.User, *domain.AppError) {
	user := new(domain.User)
	r.database.Joins("Account").First(user, "Email = ? or Account.Username = ?", username, username)
	if user.ID == "" {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) Update(u *domain.User) (*domain.User, *domain.AppError) {
	user, err := r.Find(u.ID)
	account := new(domain.Account)
	r.database.First(account, u.ID)
	if err != nil {
		return nil, err
	}
	user.Name = u.Name
	user.Email = u.Email
	account.Username = u.Account.Username
	// TODO: update password ¿?
	// user.Account.Password = u.Account.Password
	account.IsActive = u.Account.IsActive
	account.AccountType = u.Account.AccountType
	r.database.Save(user)
	r.database.Save(account)
	return user, nil
}

func (r *UserGormRepository) Delete(id string) (*domain.User, *domain.AppError) {
	user, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(user)
	return user, nil
}

func NewUserGormRepository(database *gorm.DB) domain.IUserRepository {
	database.AutoMigrate(&domain.User{})
	r := UserGormRepository{database}
	return &r
}
