package main

import (
	"fmt"
	"spectra/commom"
	"spectra/database"
	appErrors "spectra/errors"
	"spectra/routes"

	"github.com/gin-gonic/gin"
)

type Main struct {
	App *gin.Engine
}

func initConfigServer() appErrors.ErrorResponse {
	var databaseError error
	err := commom.LoadEnvs("local")

	database.Database, databaseError = database.InitDatabase()
	if databaseError != nil {
		return appErrors.InternalServerError(fmt.Sprintf("Error to connect to database - %s", databaseError.Error()))
	}
	return err
}

func main() {
	main := Main{}
	err := initConfigServer()
	defer database.Database.DisconnectDatabse()

	if err.Message != "" {
		return
	}

	main.App = gin.Default()
	main.App.MaxMultipartMemory = 15 << 20 // max 15 MiB

	v1 := main.App.Group("/api/v1/")
	{
		spectraRoutes := routes.Routes{}
		spectraRoutes.Handler(v1)
	}

	main.App.Run(commom.Envs.AppPort)
}
