package controllers

import (
	"enroll/interfaces"
	"enroll/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (authController *AuthController) ValidToken(ctx *gin.Context) {
	authService := services.AuthService{}
	var input services.ValidAccessInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, interfaces.ErrorMessage{
			Message: err.Error(),
		})
	} else {
		err := authService.ValidAccess(input.Token)
		if err.Message != "" {
			ctx.JSON(err.StatusCode(), interfaces.StringMessage{Message: err.Message})
		} else {
			ctx.Status(http.StatusOK)
		}
	}
}
