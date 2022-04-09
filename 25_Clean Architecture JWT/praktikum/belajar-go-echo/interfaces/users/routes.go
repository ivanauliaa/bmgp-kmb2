package users

import (
	"belajar-go-echo/infrastructures/http/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	echo.GET("/users", GetAllUsersHandler)
	echo.POST("/users", CreateUserHandler, middleware.JWTMiddleware())
	echo.POST("/login", LoginUser)
}
