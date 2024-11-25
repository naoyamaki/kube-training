package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/naoyamaki/usecase"
)

type PostController struct {
	postUsecase *usecase.PostUsecase
}

func NewPostController(uc *usecase.PostUsecase) *PostController {
	return &PostController{postUsecase: uc}
}

func (ctrl *PostController) RegisterPostRoutes(router *gin.Engine) {
	postGroup := router.Group("/post")
	{
		postGroup.GET("/", func(c *gin.Context) {
			statusCode, response := ctrl.postUsecase.GetAllPosts(c)
			c.JSON(statusCode, response)
		})
		postGroup.POST("/", func(c *gin.Context) {
			statusCode, response := ctrl.postUsecase.CreatePost(c)
			c.JSON(statusCode, response)
		})
		postGroup.GET("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.postUsecase.GetPostByID(c)
			c.JSON(statusCode, response)
		})
		postGroup.DELETE("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.postUsecase.DeletePost(c)
			c.JSON(statusCode, response)
		})
		postGroup.PATCH("/:id", func(c *gin.Context) {
			statusCode, response := ctrl.postUsecase.EditPost(c)
			c.JSON(statusCode, response)
		})
	}
}
