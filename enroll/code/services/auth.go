package services

import (
	"enroll/appErrors"
	"enroll/helpers"
)

type AuthService struct{}

type ValidAccessInput struct {
	Token string `json:"token" binding:"required"`
}

func (authService *AuthService) ValidAccess(token string) appErrors.ErrorResponse {
	var errorToken appErrors.ErrorResponse

	err := helpers.ValidJwtToken(token)

	if err != nil {
		errorToken = appErrors.AccessDenied(err.Error())
		return errorToken
	}

	return errorToken
}
