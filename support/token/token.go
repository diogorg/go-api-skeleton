package token

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	token := strings.Replace(auth, "Bearer ", "", 1)
	return token
}
