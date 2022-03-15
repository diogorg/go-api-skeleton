package db

import (
	"api/db/migration"
	"api/db/postgres"
	"fmt"

	"gorm.io/gorm"
)

func connect() *gorm.DB {
	return postgres.Connect()
}

func Find(dest interface{}, conds ...interface{}) {
	db := connect()
	db.Find(dest, conds...)
}

func First(dest interface{}, conds ...interface{}) {
	db := connect()
	db.First(dest, conds...)
}

func Create(value interface{}) {
	db := connect()
	db.Create(value)
}

func Delete(value interface{}, id int) {
	db := connect()
	db.Delete(value, id)
}

func Save(value interface{}) {
	db := connect()
	db.Save(value)
}

func UpdateColumns(id int, model interface{}) {
	db := connect()
	db.Model(&model).Where("id = ?", id).UpdateColumns(model)
}

func Search(where interface{}, model interface{}) {
	db := connect()
	db.Where(where).First(model)
}

func TruncateTable(table string) {
	db := connect()
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", table))
}

func TruncateTableByModel(model interface{}) {
	db := connect()
	table := db.Model(model).Name()
	db.Exec(fmt.Sprintf("TRUNCATE TABLE  %s;", table))
}

func Migrate() {
	db := connect()
	migration.Migrate(db)
}
