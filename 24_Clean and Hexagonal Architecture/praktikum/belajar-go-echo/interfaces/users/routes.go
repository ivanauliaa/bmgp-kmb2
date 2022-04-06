package users

import (
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	echo.GET("/users", GetAllUsersHandler)
	echo.POST("/users", CreateUserHandler)
}
