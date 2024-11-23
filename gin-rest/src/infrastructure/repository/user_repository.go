package repository

import (
	"fmt"

	"github.com/naoyamaki/domain/user"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetAll() ([]*user.User, error) {
	result, err := user.NewUser("asdf", "asdf@example.com", "hoge", 1)
	if err != nil {
		fmt.Println(err)
	}
	var users []*user.User
	users = append(users, result, result)
	return users, nil
}
func (r *userRepository) FindById(id string) (*user.User, error) {
	result, err := user.NewUser("asdf", "asdf@example.com", "hoge", 1)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}
func (r *userRepository) Create(user *user.User) error {
	return nil
}
func (r *userRepository) Edit(user *user.User) error {
	return nil
}
func (r *userRepository) Delete(id string) error {
	return nil
}
