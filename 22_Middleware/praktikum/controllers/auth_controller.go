package controllers

import (
	"net/http"
	"strconv"

	"praktikum/database"
	"praktikum/middleware"
	"praktikum/models"

	formatter "github.com/ivanauliaa/response-formatter"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	DB := database.Connect()

	user := models.User{}
	c.Bind(&user)

	err := DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	token, err := middleware.CreateToken(strconv.Itoa(int(user.ID)), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, formatter.InternalServerErrorResponse(map[string]interface{}{"error": err.Error()}))
	}

	responseData := models.UserLogin{}
	responseData.ID = user.ID
	responseData.Name = user.Name
	responseData.Email = user.Email
	responseData.Token = token

	return c.JSON(http.StatusOK, formatter.SuccessResponse(responseData))
}
