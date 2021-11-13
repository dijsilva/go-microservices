package routes

import (
	"spectra/controllers"
	"spectra/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (r *Routes) Handler(routerGroup *gin.RouterGroup) {
	spectraController := controllers.SpectraController{}

	routerGroup.Use(middleware.Authentication())
	routerGroup.POST("/create", spectraController.CreateSpectra)
	routerGroup.GET("/list-by-owner", spectraController.ListByOwner)
	routerGroup.GET("/spectra/:id", spectraController.GetById)

}
