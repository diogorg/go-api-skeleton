package mysql

import (
	"api/support"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	config := support.LoadConfig()
	db, err := gorm.Open(mysql.Open(config.DB_MYSQL))
	support.Panic(err)

	return db
}
