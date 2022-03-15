package services

import (
	"api/domains/auth/models"
	"api/domains/auth/repositories/providers"
	"api/domains/auth/services/passeto"
	userModel "api/domains/users/models"
)

var repository = providers.GetAuthRespository()

func Login(user models.AuthLogin) (string, error) {
	return repository.Login(user)
}

func Me(token string) userModel.User {
	payload := passeto.GetPayload(token)
	return repository.Me(payload)
}

func Verify(token string) bool {
	return passeto.Check(token)
}
