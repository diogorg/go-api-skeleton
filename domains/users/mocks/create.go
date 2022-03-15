package mocks

import (
	"api/db"
	"api/domains/users/models"
)

func CreateUser() models.User {
	user := models.User{
		Name:     "User Test",
		Email:    "user@test.com",
		Password: "12345678",
	}
	db.Create(&user)

	return user
}
