package controller

import (
	"github.com/gin-gonic/gin"

	"gin-rest/usecase"
)

type PostController struct {
	postUsecase *usecase.PostUsecase
}

func NewPostController(uc *usecase.PostUsecase) *PostController {
	return &PostController{postUsecase: uc}
}

func (ctrl *PostController) RegisterPostRoutes(router *gin.Engine) {
	router.GET("/post", func(c *gin.Context) {
		statusCode, response := ctrl.postUsecase.GetAllPosts(c)
		c.JSON(statusCode, response)
	})
	router.POST("/post", func(c *gin.Context) {
		statusCode, response := ctrl.postUsecase.CreatePost(c)
		c.JSON(statusCode, response)
	})
	router.GET("/post/:id", func(c *gin.Context) {
		statusCode, response := ctrl.postUsecase.GetPostByID(c)
		c.JSON(statusCode, response)
	})
	router.DELETE("/post/:id", func(c *gin.Context) {
		statusCode, response := ctrl.postUsecase.DeletePost(c)
		c.JSON(statusCode, response)
	})
	router.PATCH("/post/:id", func(c *gin.Context) {
		statusCode, response := ctrl.postUsecase.EditPost(c)
		c.JSON(statusCode, response)
	})
}
