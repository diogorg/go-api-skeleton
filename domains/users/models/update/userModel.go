package update

import (
	"api/support"
	"api/support/crypto"
	"time"
)

type User struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func (c *User) EncryptPassword() {
	c.Password = crypto.Hash(c.Password)
}

func (c *User) SetUpdated() {
	c.UpdatedAt = support.CurrentDateTime()
}
