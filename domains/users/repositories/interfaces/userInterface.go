package interfaces

import (
	"api/domains/users/models"
	"api/domains/users/models/update"
)

type UserRepository interface {
	GetAll() []models.User
	FindById(id int) models.User
	FindByEmail(email string) models.User
	Store(user *models.User)
	Delete(id int)
	Update(id int, user *update.User)
}
