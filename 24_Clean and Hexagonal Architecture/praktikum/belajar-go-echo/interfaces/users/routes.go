package users

import (
	"belajar-go-echo/infrastructures/database/sqlite"
	ur "belajar-go-echo/repositories/users"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	userRepo := ur.NewSQLiteRepository(db)
	userHandler := NewUserHandler(userRepo)

	echo.GET("/users", userHandler.GetUsersHandler)
	echo.POST("/users", userHandler.CreateUserHandler)
}
