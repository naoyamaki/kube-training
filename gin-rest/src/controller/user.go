package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/naoyamaki/usecase"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			statusCode, response := usecase.GetAllUsers(c)
			c.JSON(statusCode, response)
		})
		userGroup.POST("/", func(c *gin.Context) {
			statusCode, response := usecase.CreateUser(c)
			c.JSON(statusCode, response)
		})
		userGroup.GET("/:id", func(c *gin.Context) {
			statusCode, response := usecase.GetUserByID(c)
			c.JSON(statusCode, response)
		})
		userGroup.DELETE("/:id", func(c *gin.Context) {
			statusCode, response := usecase.DeleteUser(c)
			c.JSON(statusCode, response)
		})
		userGroup.PATCH("/:id", func(c *gin.Context) {
			statusCode, response := usecase.EditUser(c)
			c.JSON(statusCode, response)
		})
	}
}
