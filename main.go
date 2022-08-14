package main

import (
	"web-api/config/app"
	"web-api/config/db"
)

func main() {
	defer db.DisconnectDb()
	db.ConnectDb()
	app, server := app.Application()
	app.Listen(server)
}
