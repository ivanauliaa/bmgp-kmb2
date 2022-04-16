package users

import (
	"belajar-go-echo/infrastructures/database/sqlite"
	m "belajar-go-echo/infrastructures/http/middleware"
	ur "belajar-go-echo/repositories/users"
	us "belajar-go-echo/services/users"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	userRepo := ur.NewUserSQLiteRepository(db)
	userService := us.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	echo.GET("/users", userHandler.GetUsersHandler, m.JWTMiddleware())
	echo.POST("/users", userHandler.CreateUserHandler)
}
