package controllers

import (
	"enroll/interfaces"
	"enroll/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func (u *UserController) Create(ctx *gin.Context) {
	userService := services.UserService{}
	var input services.CreateUserInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
		})
	} else {
		error := userService.CreateUser(&input)
		if error.Message != "" {
			ctx.JSON(error.StatusCode(), interfaces.StringMessage{Message: error.Message})
		} else {
			ctx.JSON(http.StatusOK, interfaces.StringMessage{Message: "User created"})
		}
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	userService := services.UserService{}
	var input services.LoginUserInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
		})
	} else {
		user, err := userService.LoginUser(&input)
		if err.Message != "" {
			fmt.Println(err)
			ctx.JSON(err.StatusCode(), interfaces.ErrorMessage{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}
