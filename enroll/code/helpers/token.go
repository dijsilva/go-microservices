package helpers

import (
	"enroll/appErrors"
	"enroll/database/entities"
	"enroll/utils"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(user entities.User) (string, appErrors.ErrorResponse) {
	var jsonWebToken string
	var appError appErrors.ErrorResponse
	var err error
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.Id,
		"name":    user.Name,
		"email":   user.Email,
		"profile": user.Profile.ProfileName,
		"exp": time.Now().Add(time.Duration(
			utils.ConfigurationEnvs.TokenExpirationInHours,
		) * time.Hour).Unix(),
	})

	jsonWebToken, err = token.SignedString([]byte(utils.ConfigurationEnvs.TokenSignature))
	if err != nil {
		appError = appErrors.InternalServerError(err.Error())
		return jsonWebToken, appError
	}
	return jsonWebToken, appError
}

func ValidJwtToken(token string) error {
	_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Token inválido")
		}
		return []byte(utils.ConfigurationEnvs.TokenSignature), nil
	})

	return err
}
