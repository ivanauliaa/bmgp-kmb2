package users

import (
	model "belajar-go-echo/domains/users"
	"belajar-go-echo/repositories/users"

	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := users.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})

}

func GetAllUsersHandler(c echo.Context) error {
	users := users.GetAllUsers()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}
