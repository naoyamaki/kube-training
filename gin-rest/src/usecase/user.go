package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/naoyamaki/domain/user"
)

type UserUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (uc *UserUsecase) GetAllUsers(c *gin.Context) (int, gin.H) {
	users, _ := uc.userRepository.GetAll()
	return 200, gin.H{
		"id":                users[0].Id,
		"email":             users[0].Email,
		"name":              users[0].Name,
		"user_status_id_fk": users[0].User_status_id_fk,
	}
}

func (uc *UserUsecase) GetUserByID(c *gin.Context) (int, gin.H) {
	id := c.Param("id")

	var user, _ = uc.userRepository.FindById(id)
	return 200, gin.H{
		"id":                user.Id,
		"email":             user.Email,
		"name":              user.Name,
		"user_status_id_fk": user.User_status_id_fk,
	}
}

func (uc *UserUsecase) CreateUser(c *gin.Context) (int, gin.H) {
	user, _ := user.NewUser("qwer", "sample@example.com", "fuga", 2)
	result := uc.userRepository.Create(user)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Post created"}
}

func (uc *UserUsecase) DeleteUser(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	result := uc.userRepository.Delete(id)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Deleted post", "id": id}
}

func (uc *UserUsecase) EditUser(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	user, _ := user.NewUser("qwer", "sample@example.com", "fuga", 2)
	result := uc.userRepository.Edit(user)
	if result != nil {
		return 500, gin.H{"message": "error"}
	}
	return 200, gin.H{"message": "Edited post", "id": id}
}
