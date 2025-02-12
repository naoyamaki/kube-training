package controller

import (
	"github.com/gin-gonic/gin"

	"gin-rest/usecase"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

func NewUserController(uc *usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (ctrl *UserController) RegisterUserRoutes(router *gin.Engine) {
	router.GET("/user", func(c *gin.Context) {
		statusCode, response := ctrl.userUsecase.GetAllUsers(c)
		c.JSON(statusCode, response)
	})
	router.POST("/user", func(c *gin.Context) {
		statusCode, response := ctrl.userUsecase.CreateUser(c)
		c.JSON(statusCode, response)
	})
	router.GET("/user/:id", func(c *gin.Context) {
		statusCode, response := ctrl.userUsecase.GetUserByID(c)
		c.JSON(statusCode, response)
	})
	router.DELETE("/user/:id", func(c *gin.Context) {
		statusCode, response := ctrl.userUsecase.DeleteUser(c)
		c.JSON(statusCode, response)
	})
	router.PATCH("/user/:id", func(c *gin.Context) {
		statusCode, response := ctrl.userUsecase.EditUser(c)
		c.JSON(statusCode, response)
	})
}
