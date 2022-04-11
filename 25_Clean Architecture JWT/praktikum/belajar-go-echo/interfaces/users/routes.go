package users

import (
	"belajar-go-echo/infrastructures/database/sqlite"
	m "belajar-go-echo/infrastructures/http/middleware"
	ur "belajar-go-echo/repositories/users"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	userRepo := ur.NewUserSQLiteRepository(db)
	userHandler := NewUserHandler(userRepo)

	echo.GET("/users", userHandler.GetUsersHandler, m.JWTMiddleware())
	echo.POST("/users", userHandler.CreateUserHandler)
}
