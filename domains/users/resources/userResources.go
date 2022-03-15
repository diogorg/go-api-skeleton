package resources

import (
	"api/domains/users/models"
	"api/support"

	"github.com/gin-gonic/gin"
)

func GetUserResource(user models.User) gin.H {
	return gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": support.ParseDbToBr(user.CreatedAt),
		"updated_at": support.ParseDbToBr(user.UpdatedAt),
	}
}

func GetUsersResource(users []models.User) []gin.H {
	var list []gin.H
	for _, user := range users {
		list = append(list, GetUserResource(user))
	}
	return list
}
