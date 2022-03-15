package adapters

import (
	"api/domains/users/models/update"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) (update.User, error) {
	var user update.User
	err := c.ShouldBindJSON(&user)
	user.SetUpdated()
	return user, err
}
