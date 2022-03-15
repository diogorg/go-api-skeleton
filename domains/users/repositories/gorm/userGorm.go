package gorm

import (
	"api/db"
	"api/domains/users/models"
	"api/domains/users/models/update"
)

type UserRepository struct {
}

func (c UserRepository) GetAll() []models.User {
	var users []models.User
	db.Find(&users)
	return users
}

func (c UserRepository) FindById(id int) models.User {
	var user models.User
	db.First(&user, id)
	return user
}

func (c UserRepository) Store(user *models.User) {
	db.Create(user)
}

func (c UserRepository) Delete(id int) {
	var user models.User
	db.Delete(&user, id)
}

func (c UserRepository) Update(id int, user *update.User) {
	db.UpdateColumns(id, user)
}

func (c UserRepository) FindByEmail(email string) models.User {
	var user models.User
	db.Search(&models.User{Email: email}, &user)

	return user
}
