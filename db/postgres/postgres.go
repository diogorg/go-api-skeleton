package postgres

import (
	"api/support"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	config := support.LoadConfig()
	db, err := gorm.Open(postgres.Open(config.DB_POSTGRES))
	support.Panic(err)

	return db
}
