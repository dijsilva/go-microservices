package controllers

import (
	"auth-control/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controllers struct{}

type CreateTokenResponse struct {
	Data services.CreateTokenServiceResponse `json:"data"`
}

type ParseTokenResponse struct {
	Data services.ValidTokenResponse `json:"data"`
}

type ValidTokenResponse struct {
	Message string `json:"message"`
}

func (c *Controllers) CreateToken(ctx *gin.Context) {
	authService := services.Services{}
	var input services.CreateTokenInput
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenServiceResponse, err := authService.CreateToken(ctx.Request.Context(), &input)
	if err.Message != "" {
		ctx.JSON(err.StatusCode(), err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, CreateTokenResponse{Data: tokenServiceResponse})
}

func (c *Controllers) ValidToken(ctx *gin.Context) {
	authService := services.Services{}
	var input services.ValidTokenInput
	if err := ctx.ShouldBind(&input); err != nil {
		log.Println(fmt.Sprintf("Invalid fields for valid token - %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err := authService.ValidToken(ctx.Request.Context(), &input)
	if err.Message != "" {
		ctx.JSON(err.StatusCode(), ValidTokenResponse{Message: err.Message})
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Controllers) ValidParseToken(ctx *gin.Context) {
	authService := services.Services{}
	var input services.ValidTokenInput
	if err := ctx.ShouldBind(&input); err != nil {
		log.Println(fmt.Sprintf("Invalid fields for valid token - %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	tokenParsed, err := authService.ValidToken(ctx.Request.Context(), &input)
	if err.Message != "" {
		ctx.JSON(err.StatusCode(), ValidTokenResponse{Message: err.Message})
		return
	}
	ctx.JSON(http.StatusOK, ParseTokenResponse{Data: tokenParsed})
}

func (c *Controllers) DeleteToken(ctx *gin.Context) {
	authService := services.Services{}
	var input services.DeleteTokenInput
	if err := ctx.ShouldBind(&input); err != nil {
		log.Println(fmt.Sprintf("Invalid fields for delete token - %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := authService.DeleteToken(ctx.Request.Context(), &input)
	if err.Message != "" {
		ctx.JSON(err.StatusCode(), ValidTokenResponse{Message: err.Message})
		return
	}
	ctx.Status(http.StatusOK)
}
