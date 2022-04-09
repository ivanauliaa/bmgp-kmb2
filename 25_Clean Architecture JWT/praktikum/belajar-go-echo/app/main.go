package main

import (
	"belajar-go-echo/infrastructures/database/sqlite"
	"belajar-go-echo/infrastructures/http/server"
)

func init() {
	err := sqlite.MigrateDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	app := server.Server()
	app.Start(":8100")
}
