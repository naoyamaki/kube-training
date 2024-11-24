package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/naoyamaki/usecase"
)

func RegisterPostRoutes(router *gin.Engine) {
	postGroup := router.Group("/post")
	{
		postGroup.GET("/", func(c *gin.Context) {
			statusCode, response := usecase.GetAllPosts(c)
			c.JSON(statusCode, response)
		})
		postGroup.POST("/", func(c *gin.Context) {
			statusCode, response := usecase.CreatePost(c)
			c.JSON(statusCode, response)
		})
		postGroup.GET("/:id", func(c *gin.Context) {
			statusCode, response := usecase.GetPostByID(c)
			c.JSON(statusCode, response)
		})
		postGroup.DELETE("/:id", func(c *gin.Context) {
			statusCode, response := usecase.DeletePost(c)
			c.JSON(statusCode, response)
		})
		postGroup.PATCH("/:id", func(c *gin.Context) {
			statusCode, response := usecase.EditPost(c)
			c.JSON(statusCode, response)
		})
	}
}
