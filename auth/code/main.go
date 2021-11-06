package main

import (
	"auth-control/configurations"
	"auth-control/routes"

	"github.com/gin-gonic/gin"
)

type Main struct {
	appRouter *gin.Engine
}

type Message struct {
	Message string
}

func (main *Main) initConfig() error {
	var err error

	if err = configurations.LoadEnvs("local"); err != nil {
		return err
	}
	return nil
}

func main() {
	main := Main{}

	if err := main.initConfig(); err != nil {
		return
	}

	main.appRouter = gin.Default()

	v1 := main.appRouter.Group("/api/v1")
	{
		authRoute := routes.Routes{}
		authRoute.Handler(v1)
	}

	main.appRouter.Run(configurations.Envs.AppPort)
}
