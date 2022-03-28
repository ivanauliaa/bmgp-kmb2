package main

import (
	"praktikum/database"
	"praktikum/routes"
)

func init() {
	database.InitialMigration()
}

func main() {
	// create a new echo instance
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8100"))
}
