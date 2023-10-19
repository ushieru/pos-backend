package domain

type User struct {
	Model
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Account Account `json:"account"`
}

type IUserRepository interface {
	List() ([]User, *AppError)
	Save(*User) (*User, *AppError)
	Find(id uint) (*User, *AppError)
	FindByUserOrEmail(username string) (*User, *AppError)
	Update(*User) (*User, *AppError)
	Delete(id uint) (*User, *AppError)
}
