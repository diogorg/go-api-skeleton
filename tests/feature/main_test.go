package tests

import (
	"api/db"
)

func init() {
	db.Migrate()
}
