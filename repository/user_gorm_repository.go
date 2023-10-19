package repository

import (
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/utils"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	database *gorm.DB
}

func (r *UserGormRepository) seed() {
	user := new(domain.User)
	r.database.First(&user)
	if user.ID != 0 {
		return
	}
	password, _ := utils.HashPassword("admin")
	isActive := true
	r.database.Create(&domain.User{
		Name:  "Admin",
		Email: "admin@email.com",
		Account: domain.Account{
			Username:    "admin",
			Password:    password,
			IsActive:    &isActive,
			AccountType: domain.Admin,
		},
	})
	password, _ = utils.HashPassword("cashier")
	isActive = true
	r.database.Create(&domain.User{
		Name:  "Cashier",
		Email: "cashier@email.com",
		Account: domain.Account{
			Username:    "cashier",
			Password:    password,
			IsActive:    &isActive,
			AccountType: domain.Cashier,
		},
	})
	password, _ = utils.HashPassword("waiter")
	isActive = true
	r.database.Create(&domain.User{
		Name:  "Waiter",
		Email: "waiter@email.com",
		Account: domain.Account{
			Username:    "waiter",
			Password:    password,
			IsActive:    &isActive,
			AccountType: domain.Waiter,
		},
	})
}

func (r *UserGormRepository) List() ([]domain.User, *domain.AppError) {
	var user []domain.User
	r.database.Find(&user)
	return user, nil
}

func (r *UserGormRepository) Save(user *domain.User) (*domain.User, *domain.AppError) {
	result := r.database.Save(user)
	if result.RowsAffected == 0 {
		return nil, domain.NewUnexpectedError("Error al crear usuario")
	}
	return user, nil
}

func (r *UserGormRepository) Find(id uint) (*domain.User, *domain.AppError) {
	user := new(domain.User)
	r.database.First(user, id)
	if user.ID == 0 {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) FindByUserOrEmail(username string) (*domain.User, *domain.AppError) {
	user := new(domain.User)
	r.database.Joins("Account").First(user, "Email = ? or Account.Username = ?", username, username)
	if user.ID == 0 {
		return nil, domain.NewNotFoundError("Usuario no encontrado")
	}
	return user, nil
}

func (r *UserGormRepository) Update(u *domain.User) (*domain.User, *domain.AppError) {
	user, err := r.Find(u.ID)
	if err != nil {
		return nil, err
	}
	user.Name = u.Name
	user.Email = u.Email
	user.Account.Username = u.Account.Username
	// TODO: update password ¿?
	// user.Account.Password = u.Account.Password
	user.Account.IsActive = u.Account.IsActive
	user.Account.AccountType = u.Account.AccountType
	r.database.Save(user)
	return user, nil
}

func (r *UserGormRepository) Delete(id uint) (*domain.User, *domain.AppError) {
	user, err := r.Find(id)
	if err != nil {
		return nil, err
	}
	r.database.Delete(user)
	return user, nil
}

func NewUserGormRepository(database *gorm.DB) *UserGormRepository {
	database.AutoMigrate(&domain.User{})
	r := UserGormRepository{database}
	r.seed()
	return &r
}
