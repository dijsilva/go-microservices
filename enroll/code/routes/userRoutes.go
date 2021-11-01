package routes

import (
	"enroll/controllers"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) Handler(routerGroup *gin.RouterGroup) {
	uC := controllers.UserController{}
	routerGroup.POST("/new", uC.Create)
	routerGroup.POST("/auth", uC.Login)
}
