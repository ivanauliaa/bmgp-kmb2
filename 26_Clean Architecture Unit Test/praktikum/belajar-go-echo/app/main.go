package main

import (
	"belajar-go-echo/infrastructures/http/server"
)

func main() {
	app := server.Server()
	app.Start(":8100")
}
