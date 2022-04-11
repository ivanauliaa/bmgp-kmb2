package auth

import (
	"belajar-go-echo/domains/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type handler struct {
	repository auth.AuthRepository
}

func NewAuthHandler(repo auth.AuthRepository) auth.AuthHandler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) LoginHandler(c echo.Context) error {
	credential := auth.Auth{}
	err := c.Bind(&credential)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.repository.Login(credential)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   jwt,
	})

}
