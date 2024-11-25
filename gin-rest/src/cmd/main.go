package main

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/naoyamaki/controller"
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

	userController := controller.NewUserController()
	postController := controller.NewPostController()
	// TODO: infrastracture importしてここでnewする
	// db.Newpostgres(conf.DB)
	router := gin.Default()
	userController.RegisterUserRoutes(router)
	postController.RegisterPostRoutes(router)

	fmt.Println("start up " + cfg.Host + cfg.Port)
	router.Run(cfg.Port)
}
