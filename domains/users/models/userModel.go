package models

import (
	"api/support/crypto"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"nonzero"`
	Email    string `json:"email" validate:"nonzero,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"nonzero,min=8"`
}

func (c *User) EncryptPassword() {
	c.Password = crypto.Hash(c.Password)
}

func (c *User) Validate() error {
	err := validator.Validate(c)
	if err != nil {
		return err
	}

	return nil
}
