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

	// TODO: infrastracture importしてここでnewする
	// db.Newpostgres(conf.DB)
	router := gin.Default()
	controller.RegisterUserRoutes(router)
	controller.RegisterPostRoutes(router)

	fmt.Println("start up " + cfg.Host + cfg.Port)
	router.Run(cfg.Port)
}
