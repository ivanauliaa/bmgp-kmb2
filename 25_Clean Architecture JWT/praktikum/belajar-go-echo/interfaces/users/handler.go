package users

import (
	model "belajar-go-echo/domains/users"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	repository model.UserRepository
}

func NewUserHandler(repo model.UserRepository) model.UserHandler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) CreateUserHandler(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	err := h.repository.CreateUser(user)
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
	users := h.repository.GetUsers()

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

	user, err := h.repository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"user":     user,
	})
}
