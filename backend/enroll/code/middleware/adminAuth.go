package middleware

import (
	"enroll/interfaces"
	"enroll/providers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data: interfaces.ErrorMessage{
					Message: "Token not provided",
					Status:  401,
				},
			})
			return
		}

		bearerAndToken := strings.Fields(authHeader)

		if len(bearerAndToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data: interfaces.ErrorMessage{
					Message: "Token bad formated",
					Status:  401,
				},
			})
			return
		}

		authControl := providers.AuthControl{}
		respValidToken, err := authControl.ValidToken(providers.ValidTokenInput{
			Token:     bearerAndToken[1],
			TokenKind: "LOGIN_ADMIN",
		})

		if err.Message != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data: interfaces.ErrorMessage{
					Message: err.Message,
					Status:  http.StatusUnauthorized,
				},
			})
			return

		}
		if !respValidToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, interfaces.ErrorResponse{
				Data: interfaces.ErrorMessage{
					Message: "Access not allowed to this endpoint",
					Status:  http.StatusUnauthorized,
				},
			})
			return
		}

		c.Next()
	}
}
