package services

import (
	appErrors "auth-control/errors"
)

type Services struct{}

type CreateTokenInput struct {
	UserId    string `json:"userId" binding:"required"`
	Profile   string `json:"profile" binding:"required"`
	UserMail  string `json:"userMail" binding:"required"`
	TokenKind string `json:"tokenKind" binding:"required"`
}

func (s *Services) CreateToken(input *CreateTokenInput) appErrors.ErrorResponse {
	err := appErrors.Unauthorized("Not authorized")
	return err
}
