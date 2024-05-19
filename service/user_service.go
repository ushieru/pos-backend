package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
	"github.com/ushieru/pos/utils"
)

type IUserService interface {
	List() ([]domain.User, *domain.AppError)
	Find(id uint) (*domain.User, *domain.AppError)
	AuthWithCredentials(username, password, secret string) (*dto.AuthUserResponse, *domain.AppError)
	Save(dto *dto.CreateUserRequest, a *domain.Account) (*domain.User, *domain.AppError)
	Update(id uint, dto *dto.UpdateUserRequest, a *domain.Account) (*domain.User, *domain.AppError)
	Delete(id uint, a *domain.Account) (*domain.User, *domain.AppError)
}

type UserService struct {
	repository domain.IUserRepository
}

func (c *UserService) List() ([]domain.User, *domain.AppError) {
	return c.repository.List()
}

func (c *UserService) Find(id uint) (*domain.User, *domain.AppError) {
	return c.repository.Find(id)
}

func (c *UserService) AuthWithCredentials(username, password, secret string) (*dto.AuthUserResponse, *domain.AppError) {
	user, err := c.repository.FindByUserOrEmail(username)
	if err != nil {
		return nil, err
	}
	matchPassword := utils.CheckPasswordHash(password, user.Account.Password)
	if !matchPassword {
		return nil, domain.NewNotFoundError("Credenciales incorrectas")
	}
	if !*user.Account.IsActive {
		return nil, domain.NewUnauthorizedError("Usuario desactivado consulte al administrador")
	}
	claims := jwt.MapClaims{
		"SessionParamAdminId": user.Account.ID,
		"SessionParamUserId":  user.ID,
		"SessionParamRole":    string(user.Account.AccountType),
		"exp":                 time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	authUserResponse := &dto.AuthUserResponse{
		Token: tokenString,
		User:  *user,
	}
	return authUserResponse, nil
}

func (c *UserService) Save(dto *dto.CreateUserRequest, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	isActive := true
	hashPassword, _ := utils.HashPassword(dto.Password)
	user := &domain.User{
		Name:  dto.Name,
		Email: dto.Email,
		Account: domain.Account{
			Username:    dto.Username,
			Password:    hashPassword,
			IsActive:    &isActive,
			AccountType: domain.AccountType(dto.AccountType),
		},
	}
	return c.repository.Save(user)
}

func (c *UserService) Update(id uint, dto *dto.UpdateUserRequest, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	if id == a.ID && !dto.IsActive {
		return nil, domain.NewUnauthorizedError("No puedes desactivarte a ti mismo")
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	// TODO: update password ¿?
	hashPassword, _ := utils.HashPassword(dto.Password)
	user := &domain.User{
		Model: domain.Model{
			ID: id,
		},
		Name:  dto.Name,
		Email: dto.Email,
		Account: domain.Account{
			Username:    dto.Username,
			Password:    hashPassword,
			IsActive:    &dto.IsActive,
			AccountType: domain.AccountType(dto.AccountType),
		},
	}
	return c.repository.Update(user)
}

func (c *UserService) Delete(id uint, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	return c.repository.Delete(id)
}

func NewUserService(repository domain.IUserRepository) *UserService {
	return &UserService{repository}
}
