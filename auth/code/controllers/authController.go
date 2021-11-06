package controllers

import (
	"auth-control/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controllers struct{}

func (c *Controllers) CreateToken(ctx *gin.Context) {
	//authService := services.Services{}
	var input services.CreateTokenInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	// authService.CreateToken()
}
