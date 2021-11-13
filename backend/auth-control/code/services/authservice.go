package services

import (
	"auth-control/configurations"
	"auth-control/database"
	appErrors "auth-control/errors"
	"context"
	"fmt"
	"log"
	"time"
)

type Services struct{}

type CreateTokenServiceResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type CreateTokenInput struct {
	UserId    string `json:"userId" binding:"required"`
	Profile   string `json:"profile" binding:"required"`
	UserMail  string `json:"userMail" binding:"required"`
	UserName  string `json:"userName" binding:"required"`
	TokenKind string `json:"tokenKind" binding:"required"`
}

type ValidTokenInput struct {
	Token     string `json:"token" binding:"required"`
	TokenKind string `json:"tokenKind" binding:"required"`
}

type DeleteTokenInput struct {
	Token string `json:"token" binding:"required"`
}

func (s *Services) CreateToken(ctx context.Context, input *CreateTokenInput) (CreateTokenServiceResponse, appErrors.ErrorResponse) {
	log.Println("Generating token")
	token, err := GenerateJwtToken(input)
	if err.Message != "" {
		return CreateTokenServiceResponse{}, err
	}

	expirationToken := time.Hour * time.Duration(configurations.Envs.TokenExpirationInHours)

	log.Println("Storing token at database")
	key, err := database.Database.Set(ctx, input.UserId, token.token, expirationToken)
	if err.Message != "" {
		return CreateTokenServiceResponse{}, err
	}
	log.Println(fmt.Sprintf("Key %s stored", key))
	return CreateTokenServiceResponse{Token: token.token, ExpiresAt: token.tokenExpiration}, appErrors.ErrorResponse{}
}

func (s *Services) ValidToken(ctx context.Context, input *ValidTokenInput) (ValidTokenResponse, appErrors.ErrorResponse) {
	log.Println("Decoding token")
	tokenValidResponse, err := DecodeJwtToken(input.Token)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to decode token - %s", err.Message))
		return ValidTokenResponse{}, err
	}

	log.Println(fmt.Sprintf("Getting token for id - %s", tokenValidResponse.Id))
	tokenStored, err := database.Database.Get(ctx, tokenValidResponse.Id)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to get token from database - %s", err.Message))
		return ValidTokenResponse{}, appErrors.Unauthorized("invalid token")
	}

	log.Println("Decoding token stored in database")
	tokenStoredParsed, err := DecodeJwtToken(tokenStored)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to decode token stored in database - %s", err.Message))
		return ValidTokenResponse{}, err
	}
	if tokenStoredParsed.TokenKind != input.TokenKind || tokenStored != input.Token {
		err = appErrors.Unauthorized("invalid token kind")
		return ValidTokenResponse{}, err
	}

	return tokenStoredParsed, appErrors.ErrorResponse{}
}

func (s *Services) DeleteToken(ctx context.Context, input *DeleteTokenInput) appErrors.ErrorResponse {
	log.Println("Decoding token to delete")
	tokenDecoded, err := DecodeJwtToken(input.Token)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to decode token - %s", err.Message))
		return err
	}

	tokenStored, err := database.Database.Get(ctx, tokenDecoded.Id)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to get token from database - %s", err.Message))
		return appErrors.ErrorResponse{}
	}

	if tokenStored != input.Token {
		log.Println("Token not valid")
		return appErrors.Unauthorized("invalid token")
	}

	log.Println(fmt.Sprintf("Getting token for id - %s", tokenDecoded.Id))
	err = database.Database.Delete(ctx, tokenDecoded.Id)
	if err.Message != "" {
		log.Println(fmt.Sprintf("Error to get token from database - %s", err.Message))
	}

	return appErrors.ErrorResponse{}
}
