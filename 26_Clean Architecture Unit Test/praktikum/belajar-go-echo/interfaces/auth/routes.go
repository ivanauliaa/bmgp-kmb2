package auth

import (
	"belajar-go-echo/infrastructures/database/sqlite"
	ar "belajar-go-echo/repositories/auth"
	as "belajar-go-echo/services/auth"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	authRepo := ar.NewAuthSQLiteRepository(db)
	authService := as.NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService)

	echo.POST("/auth/login", authHandler.LoginHandler)
}
