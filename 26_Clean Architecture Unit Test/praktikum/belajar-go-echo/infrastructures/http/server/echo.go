package server

import (
	"belajar-go-echo/interfaces/auth"
	"belajar-go-echo/interfaces/users"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()

	users.Routes(app)
	auth.Routes(app)

	return app
}
