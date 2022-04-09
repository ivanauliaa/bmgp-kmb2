package users

import (
	model "belajar-go-echo/domains/users"
	"belajar-go-echo/repositories/users"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt"
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

func LoginUser(c echo.Context) error {
	request := model.LoginRequest{}

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	user := users.Login(request)
	if user == nil {
		return c.JSON(http.StatusBadRequest, "Invalid credentials")
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, jwt)
}
