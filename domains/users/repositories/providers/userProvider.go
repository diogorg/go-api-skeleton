package providers

import (
	user "api/domains/users/repositories/gorm"
	interfaces "api/domains/users/repositories/interfaces"
)

func GetUserRespository() interfaces.UserRepository {
	repository := user.UserRepository{}

	return repository
}
