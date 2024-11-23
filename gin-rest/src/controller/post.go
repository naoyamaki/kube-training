package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.Engine) {
	postGroup := router.Group("/post")
	{
		postGroup.GET("/", getAllPosts)
		postGroup.POST("/", createPost)
		postGroup.GET("/:id", getPostByID)
		postGroup.DELETE("/:id", deletePost)
		postGroup.PATCH("/:id", editPost)
	}
}

func getAllPosts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all posts"})
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Get post by ID", "id": id})
}
func createPost(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Post created"})
}
func deletePost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Deleted post", "id": id})
}
func editPost(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"message": "Edited post", "id": id})
}
