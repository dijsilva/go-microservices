package routes

import (
	"auth-control/controllers"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (r *Routes) Handler(routerGroup *gin.RouterGroup) {
	authControlelr := controllers.Controllers{}
	routerGroup.POST("/create-token", authControlelr.CreateToken)
}
