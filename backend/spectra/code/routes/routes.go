package routes

import (
	"spectra/controllers"
	"spectra/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (r *Routes) Handler(routerGroup *gin.RouterGroup) {
	spectraController := controllers.SpectraController{}

	routerGroup.POST("/prediction/:id", spectraController.UpdatePrediction)
	routerGroup.POST("/create", middleware.Authentication(), spectraController.CreateSpectra)
	routerGroup.GET("/list-by-owner", middleware.Authentication(), spectraController.ListByOwner)
	routerGroup.GET("/spectra/:id", spectraController.GetById)

}
