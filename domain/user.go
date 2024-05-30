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
	Find(id string) (*User, *AppError)
	FindByUserOrEmailAndPassword(username, password string) (*User, *AppError)
	Update(*User) (*User, *AppError)
	Delete(*User) (*User, *AppError)
}
