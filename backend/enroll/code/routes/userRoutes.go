package routes

import (
	"enroll/controllers"
	"enroll/middleware"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) Handler(routerGroup *gin.RouterGroup) {
	uC := controllers.UserController{}
	routerGroup.POST("/new", uC.Create)
	routerGroup.POST("/auth", uC.Login)

	routerGroup.Use(middleware.AdminAuth())
	routerGroup.GET("/users", uC.ListUsers)
	routerGroup.POST("/change-profile", uC.ChangeProfile)
}
