package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ushieru/pos/domain"
	"github.com/ushieru/pos/dto"
)

type IUserService interface {
	List() ([]domain.User, *domain.AppError)
	Find(id string) (*domain.User, *domain.AppError)
	AuthWithCredentials(username, password, secret string) (*dto.AuthUserResponse, *domain.AppError)
	Save(dto *dto.CreateUserRequest, a *domain.Account) (*domain.User, *domain.AppError)
	Update(id string, dto *dto.UpdateUserRequest, a *domain.Account) (*domain.User, *domain.AppError)
	Delete(id string, a *domain.Account) (*domain.User, *domain.AppError)
}

type UserService struct {
	repository domain.IUserRepository
}

func (s *UserService) List() ([]domain.User, *domain.AppError) {
	return s.repository.List()
}

func (s *UserService) Find(id string) (*domain.User, *domain.AppError) {
	return s.repository.Find(id)
}

func (s *UserService) AuthWithCredentials(username, password, secret string) (*dto.AuthUserResponse, *domain.AppError) {
	h := sha256.New()
	h.Write([]byte(password))
	hashPassword := hex.EncodeToString(h.Sum(nil))
	user, err := s.repository.FindByUserOrEmailAndPassword(username, hashPassword)
	if err != nil {
		return nil, err
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

func (s *UserService) Save(dto *dto.CreateUserRequest, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	isActive := true
	h := sha256.New()
	h.Write([]byte(dto.Password))
	hashPassword := hex.EncodeToString(h.Sum(nil))
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
	return s.repository.Save(user)
}

func (s *UserService) Update(id string, dto *dto.UpdateUserRequest, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	if id == a.ID && !dto.IsActive {
		return nil, domain.NewUnauthorizedError("No puedes desactivarte a ti mismo")
	}
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	user, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	user.Name = dto.Name
	user.Email = dto.Email
	user.Account.Username = dto.Username
	user.Account.IsActive = &dto.IsActive
	user.Account.AccountType = domain.AccountType(dto.AccountType)
	return s.repository.Update(user)
}

func (s *UserService) Delete(id string, a *domain.Account) (*domain.User, *domain.AppError) {
	if a.AccountType != domain.Admin {
		return nil, domain.NewUnauthorizedError("No tienes autorizacion para esta accion")
	}
	if id == a.ID {
		return nil, domain.NewUnauthorizedError("No puedes eliminarte a ti mismo")
	}
	user, err := s.Find(id)
	if err != nil {
		return nil, err
	}
	return s.repository.Delete(user)
}

func (s *UserService) seed() {
	users, err := s.List()
	if err != nil {
		return
	}
	if len(users) != 0 {
		return
	}
	adminMooc := &domain.Account{AccountType: domain.Admin}
	dtoUsers := []dto.CreateUserRequest{
		{
			Name:        "admin",
			Email:       "admin@email.com",
			Username:    "admin",
			Password:    "admin",
			AccountType: string(domain.Admin),
		},
		{
			Name:        "cashier",
			Email:       "cashier@email.com",
			Username:    "cashier",
			Password:    "cashier",
			AccountType: string(domain.Cashier),
		},
		{
			Name:        "waiter",
			Email:       "waiter@email.com",
			Username:    "waiter",
			Password:    "waiter",
			AccountType: string(domain.Waiter),
		},
		{
			Name:        "cook",
			Email:       "cook@email.com",
			Username:    "cook",
			Password:    "cook",
			AccountType: string(domain.Producer),
		},
		{
			Name:        "bartender",
			Email:       "bartender@email.com",
			Username:    "bartender",
			Password:    "bartender",
			AccountType: string(domain.Producer),
		},
	}

	for _, dto := range dtoUsers {
		s.Save(&dto, adminMooc)
	}
}

func NewUserService(repository domain.IUserRepository) *UserService {
	service := &UserService{repository}
	service.seed()
	return service
}
