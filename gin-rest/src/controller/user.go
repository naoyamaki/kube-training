package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/naoyamaki/usecase"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

func NewUserController(uc *usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (ctrl *UserController) RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			statusCode, response := ctrl.userUsecase.GetAllUsers(c)
			c.JSON(statusCode, response)
		})
		userGroup.POST("/", func(c *gin.Context) {
			statusCode, response := ctrl.userUsecase.CreateUser(c)
			c.JSON(statusCode, response)
		})
		userGroup.GET("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.userUsecase.GetUserByID(c)
			c.JSON(statusCode, response)
		})
		userGroup.DELETE("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.userUsecase.DeleteUser(c)
			c.JSON(statusCode, response)
		})
		userGroup.PATCH("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.userUsecase.EditUser(c)
			c.JSON(statusCode, response)
		})
	}
}
