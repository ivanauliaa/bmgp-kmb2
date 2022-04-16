package auth

import (
	"belajar-go-echo/domains/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service auth.AuthService
}

func NewAuthHandler(serv auth.AuthService) auth.AuthHandler {
	return &handler{
		service: serv,
	}
}

func (h *handler) LoginHandler(c echo.Context) error {
	credential := auth.Auth{}
	err := c.Bind(&credential)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	token, statusCode := h.service.Login(credential)
	switch statusCode {
	case http.StatusBadRequest:
		return c.JSON(http.StatusBadRequest, "Invalid email or password")
	case http.StatusInternalServerError:
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   *token,
	})

}
