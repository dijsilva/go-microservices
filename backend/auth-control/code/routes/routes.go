package routes

import (
	"auth-control/controllers"

	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (r *Routes) Handler(routerGroup *gin.RouterGroup) {
	authControlelr := controllers.Controllers{}
	routerGroup.POST("/create-token", authControlelr.CreateToken)
	routerGroup.POST("/valid-token", authControlelr.ValidToken)
	routerGroup.POST("/valid-parse-token", authControlelr.ValidParseToken)
	routerGroup.DELETE("/logout", authControlelr.DeleteToken)
}
