package middleware

import (
	"enroll/appErrors"
	"enroll/providers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithError(http.StatusUnauthorized, appErrors.Unauthorized("Token not provided"))
		}

		bearerAndToken := strings.Fields(authHeader)

		if len(bearerAndToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, appErrors.Unauthorized("Token bad formated"))
		}

		authControl := providers.AuthControl{}
		respValidToken, err := authControl.ValidToken(providers.ValidTokenInput{
			Token:     bearerAndToken[1],
			TokenKind: "LOGIN_ADMIN",
		})

		if err.Message != "" {
			c.AbortWithStatusJSON(err.StatusCode(), err.Message)

		}
		if !respValidToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Access not allowed to this endpoint")

		}

		c.Next()
	}
}
