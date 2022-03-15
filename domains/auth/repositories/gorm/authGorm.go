package gorm

import (
	"api/db"
	"api/domains/auth/models"
	"api/domains/auth/services/passeto"
	userModel "api/domains/users/models"
	"api/support/crypto"
	"errors"
)

type AuthRepository struct {
}

func (c AuthRepository) Login(user models.AuthLogin) (string, error) {
	var userFound userModel.User
	password := crypto.Hash(user.Password)
	db.Search(&userModel.User{Email: user.Email, Password: password}, &userFound)
	if userFound.ID == 0 {
		return "", errors.New("User not found")
	}

	return passeto.Create(user.Email), nil
}

func (c AuthRepository) Me(payload *passeto.Payload) userModel.User {
	var user userModel.User
	db.Search(&userModel.User{Email: payload.Username}, &user)

	return user
}
