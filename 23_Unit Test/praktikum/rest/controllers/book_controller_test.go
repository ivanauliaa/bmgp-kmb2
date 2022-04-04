package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple/database"
	"simple/models"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreateBook(title, year, author string) *models.Book {
	DB := database.Connect()
	book := models.Book{
		Title:  title,
		Year:   year,
		Author: author,
	}

	if err := DB.Create(&book).Error; err != nil {
		return nil
	}

	return &book
}

func TruncateBooksTable() {
	DB := database.Connect()
	DB.Exec("TRUNCATE TABLE books")
}

func TestCreateBookController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create book normal",
			path:       "/restricted/books",
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"title":  "title1",
		"year":   "1997",
		"author": "author1",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Book models.Book `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.Book.Title, "title1")
			assert.Equal(t, response.Book.Year, "1997")
			assert.Equal(t, response.Book.Author, "author1")
		}
	}

	TruncateUsersTable()
	TruncateBooksTable()
}

func TestGetBooksController(t *testing.T) {
	database.InitialMigration()

	books := []models.Book{
		{
			Title: "title1", Year: "1997", Author: "author1",
		},
		{
			Title: "title2", Year: "1997", Author: "author2",
		},
	}

	for _, book := range books {
		CreateBook(book.Title, book.Year, book.Author)
	}

	testCases := []struct {
		name       string
		path       string
		expectCode int
		dataSize   int
	}{
		{
			name:       "get all books normal",
			path:       "/books",
			expectCode: http.StatusOK,
			dataSize:   len(books),
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Books []models.Book `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCase.dataSize, len(response.Books))
		}
	}

	TruncateBooksTable()
}

func TestGetBookController(t *testing.T) {
	database.InitialMigration()

	book := CreateBook("title1", "1997", "author1")

	testCases := []struct {
		name       string
		path       string
		book_id    string
		expectCode int
	}{
		{
			name:       "get book normal",
			path:       "/books",
			book_id:    strconv.Itoa(int(book.ID)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.book_id)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Book models.Book `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.Book.Title, "title1")
			assert.Equal(t, response.Book.Year, "1997")
			assert.Equal(t, response.Book.Author, "author1")
		}
	}

	TruncateBooksTable()
}

func TestUpdateBookController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()
	book := CreateBook("title1", "1997", "author1")

	testCases := []struct {
		name       string
		path       string
		book_id    string
		expectCode int
	}{
		{
			name:       "update book normal",
			path:       "/restricted/books",
			book_id:    strconv.Itoa(int(book.ID)),
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"title":  "updated_title",
		"year":   "2000",
		"author": "updated_author",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/restricted/books", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.book_id)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Book models.Book `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.Book.Title, "updated_title")
			assert.Equal(t, response.Book.Year, "2000")
			assert.Equal(t, response.Book.Author, "updated_author")
		}
	}

	TruncateUsersTable()
	TruncateBooksTable()
}

func TestDeleteBookController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()
	book := CreateBook("title1", "1997", "author1")

	testCases := []struct {
		name       string
		path       string
		book_id    string
		expectCode int
	}{
		{
			name:       "delete book normal",
			path:       "/restricted/books",
			book_id:    strconv.Itoa(int(book.ID)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/restricted/books", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.book_id)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Book interface{} `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Nil(t, response.Book)
		}
	}

	TruncateUsersTable()
	TruncateBooksTable()
}
