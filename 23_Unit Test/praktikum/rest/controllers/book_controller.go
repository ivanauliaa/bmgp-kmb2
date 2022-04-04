package controllers

import (
	"net/http"
	"simple/database"
	"simple/models"

	formatter "github.com/ivanauliaa/response-formatter"
	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	DB := database.Connect()
	var books []models.Book

	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, formatter.SuccessResponse(books))
}

func GetBookController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")
	book := models.Book{}

	DB.Where("ID = ?", id).First(&book)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
	}

	return c.JSON(http.StatusOK, formatter.SuccessResponse(book))
}

func CreateBookController(c echo.Context) error {
	DB := database.Connect()
	book := models.Book{}
	c.Bind(&book)

	if err := DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, formatter.SuccessResponse(book))
}

func DeleteBookController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")

	DB.Delete(&models.Book{}, id)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(nil))
}

func UpdateBookController(c echo.Context) error {
	DB := database.Connect()
	id := c.Param("id")
	book := models.Book{}

	DB.Where("ID = ?", id).First(&book)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, formatter.NotFoundResponse(nil))
	}

	payload := models.Book{}
	c.Bind(&payload)

	book.Title = payload.Title
	book.Author = payload.Author
	book.Year = payload.Year
	DB.Save(&book)

	return c.JSON(http.StatusOK, formatter.SuccessResponse(book))
}
