package usecase

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) (int, gin.H) {
	return 200, gin.H{"message": "Get all posts"}
}
func GetUserByID(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Get post by ID", "id": id}
}
func CreateUser(c *gin.Context) (int, gin.H) {
	return 200, gin.H{"message": "Post created"}
}
func DeleteUser(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Deleted post", "id": id}
}
func EditUser(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Edited post", "id": id}
}
