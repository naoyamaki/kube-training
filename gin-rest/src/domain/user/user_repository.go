package user

type UserRepository interface {
	GetAll() ([]*User, error)
	FindById(id string) (*User, error)
	Create(user *User) error
	Edit(user *User) error
	Delete(id string) error
}
