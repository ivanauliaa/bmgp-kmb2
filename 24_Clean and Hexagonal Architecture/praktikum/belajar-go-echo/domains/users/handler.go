package users

import "github.com/labstack/echo/v4"

type UserHandler interface {
	CreateUserHandler(c echo.Context) error
	GetUsersHandler(c echo.Context) error
}
