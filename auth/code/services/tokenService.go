package services

import (
	"auth-control/configurations"
	appErrors "auth-control/errors"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GenerateJWTTokenResponse struct {
	token           string
	tokenExpiration int64
}

type ValidTokenResponse struct {
	Email     string
	Id        string
	Name      string
	TokenKind string
	Profile   string
}

type TokenClaims struct {
	Email     string
	Id        string
	Name      string
	Profile   string
	TokenKind string
	jwt.StandardClaims
}

func GenerateJwtToken(input *CreateTokenInput) (GenerateJWTTokenResponse, appErrors.ErrorResponse) {
	var jsonWebToken string
	var appError appErrors.ErrorResponse
	var err error
	tokenExpiration := time.Now().Add(time.Duration(
		configurations.Envs.TokenExpirationInHours,
	) * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        input.UserId,
		"name":      input.UserName,
		"email":     input.UserMail,
		"tokenKind": input.TokenKind,
		"profile":   input.Profile,
		"exp":       tokenExpiration,
	})

	jsonWebToken, err = token.SignedString([]byte(configurations.Envs.TokenSignature))
	if err != nil {
		appError = appErrors.InternalServerError(err.Error())
		return GenerateJWTTokenResponse{}, appError
	}
	return GenerateJWTTokenResponse{token: jsonWebToken, tokenExpiration: tokenExpiration}, appError
}

func DecodeJwtToken(token string) (ValidTokenResponse, appErrors.ErrorResponse) {
	tokenClaims := &TokenClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(configurations.Envs.TokenSignature), nil
	})

	if err != nil {
		return ValidTokenResponse{}, appErrors.Unauthorized("invalid token")
	}
	return ValidTokenResponse{
		Email:     tokenClaims.Email,
		Id:        tokenClaims.Id,
		Name:      tokenClaims.Name,
		Profile:   tokenClaims.Profile,
		TokenKind: tokenClaims.TokenKind,
	}, appErrors.ErrorResponse{}
}
