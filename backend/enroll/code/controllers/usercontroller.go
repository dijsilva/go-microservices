package controllers

import (
	"enroll/interfaces"
	"enroll/services"
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
			ctx.JSON(error.StatusCode(), interfaces.Response{Data: interfaces.Message{Message: error.Message}})
		} else {
			ctx.JSON(http.StatusCreated, interfaces.Response{Data: interfaces.Message{Message: "User created"}})
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
			ctx.JSON(err.StatusCode(), interfaces.ErrorMessage{Message: err.Error()})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}

func (u *UserController) ListUsers(ctx *gin.Context) {
	userService := services.UserService{}
	users := userService.ListUsers()
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) ChangeProfile(ctx *gin.Context) {
	userService := services.UserService{}
	var input services.ChangeProfileInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
		})
		return
	}

	err := userService.ChangeProfile(input)
	if err.Message != "" {
		ctx.AbortWithStatusJSON(err.StatusCode(), interfaces.ErrorResponse{
			Data: interfaces.ErrorMessage{
				Message: err.Message,
				Status:  err.StatusCode(),
			},
		})
		return
	}
	ctx.Status(http.StatusOK)
}
