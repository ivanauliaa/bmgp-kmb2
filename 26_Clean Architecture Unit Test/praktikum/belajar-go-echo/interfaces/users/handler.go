package users

import (
	"belajar-go-echo/domains/users"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service users.UserService
}

func NewUserHandler(serv users.UserService) users.UserHandler {
	return &handler{
		service: serv,
	}
}

func (h *handler) CreateUserHandler(c echo.Context) error {
	user := users.User{}
	c.Bind(&user)

	err := h.service.CreateUserService(user)
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

func (h *handler) GetUsersHandler(c echo.Context) error {
	users := h.service.GetUsersService()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	})
}

func (h *handler) GetUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.GetUserService(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"user":     user,
	})
}
