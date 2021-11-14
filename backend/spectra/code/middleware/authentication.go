package middleware

import (
	"log"
	"net/http"
	"spectra/interfaces"
	"spectra/providers"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Authentication middleware")
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data:   "Token not provided",
				Status: http.StatusUnauthorized,
			})
			return
		}

		bearerAndToken := strings.Fields(authHeader)

		if len(bearerAndToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data:   "Token bad formated",
				Status: http.StatusUnauthorized,
			})
			return
		}

		authControl := providers.AuthControl{}
		authControlResponse, err := authControl.ValidToken(providers.ValidTokenInput{
			Token:     bearerAndToken[1],
			TokenKind: "LOGIN_USER",
		})

		if err.Message != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data:   err.Message,
				Status: err.StatusCode(),
			})
			return
		}

		c.Set("user_owner_email", authControlResponse.Data.Email)
	}
}
