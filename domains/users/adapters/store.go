package adapters

import (
	"api/domains/users/models"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) (models.User, error) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	invalid := user.Validate()
	if invalid != nil {
		err = invalid
	}
	user.EncryptPassword()

	return user, err
}
