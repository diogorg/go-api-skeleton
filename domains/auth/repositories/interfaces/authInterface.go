package interfaces

import (
	"api/domains/auth/models"
	"api/domains/auth/services/passeto"
	userModel "api/domains/users/models"
)

type AuthRepository interface {
	Login(user models.AuthLogin) (string, error)
	Me(payload *passeto.Payload) userModel.User
}
