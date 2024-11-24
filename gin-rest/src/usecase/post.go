package usecase

import "github.com/gin-gonic/gin"

func GetAllPosts(c *gin.Context) (int, gin.H) {
	return 200, gin.H{"message": "Get all posts"}
}

func GetPostByID(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Get post by ID", "id": id}
}
func CreatePost(c *gin.Context) (int, gin.H) {
	return 200, gin.H{"message": "Post created"}
}
func DeletePost(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Deleted post", "id": id}
}
func EditPost(c *gin.Context) (int, gin.H) {
	id := c.Param("id")
	return 200, gin.H{"message": "Edited post", "id": id}
}
