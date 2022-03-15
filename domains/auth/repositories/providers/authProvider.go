package providers

import (
	auth "api/domains/auth/repositories/gorm"
	interfaces "api/domains/auth/repositories/interfaces"
)

func GetAuthRespository() interfaces.AuthRepository {
	repository := auth.AuthRepository{}

	return repository
}
