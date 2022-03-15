package middleware

import (
	"api/domains/auth/services"
	"api/support/response"
	"api/support/token"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	token := token.GetToken(c)
	if !services.Verify(token) {
		response.JsonUnauthorized(c)
	} else {
		c.Next()
	}
}
