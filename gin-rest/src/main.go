package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/naoyamaki/controller"
	"github.com/naoyamaki/infrastructure/repository"
	"github.com/naoyamaki/usecase"
)

func main() {
	type config struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port string `env:"PORT" envDefault:":8081"`
	}
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}
	// TODO: infrastracture importしてここでnewする
	// db.Newpostgres(conf.DB)

	userRepository := repository.NewUserRepository()
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	postRepository := repository.NewPostRepository()
	postUsecase := usecase.NewPostUsecase(postRepository)
	postController := controller.NewPostController(postUsecase)

	router := gin.Default()
	userController.RegisterUserRoutes(router)
	postController.RegisterPostRoutes(router)

	fmt.Println("start up " + cfg.Host + cfg.Port)
	router.Run(cfg.Port)
}
