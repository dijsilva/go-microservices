package routes

import (
	"enroll/controllers"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func (auth *Auth) Handler(routerGroup *gin.RouterGroup) {
	aC := controllers.AuthController{}
	routerGroup.POST("/valid", aC.ValidToken)
}
