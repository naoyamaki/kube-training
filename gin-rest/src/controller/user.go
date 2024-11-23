package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", getAllUsers)
		userGroup.POST("/", createUser)
		userGroup.GET("/:id", getUserByID)
		userGroup.DELETE("/:id", deleteUser)
		userGroup.PATCH("/:id", editUser)
	}
}

func getAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all posts"})
}
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get post by ID", "id": id})
}
func createUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Post created"})
}
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Deleted post", "id": id})
}
func editUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Edited post", "id": id})
}
