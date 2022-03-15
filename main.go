package main

import (
	"api/db"
	"api/server"
)

func main() {
	db.Migrate()
	server.Run()
}
