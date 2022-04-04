package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"simple/database"
	"simple/models"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func CreateUser(name, email, password string) *models.User {
	DB := database.Connect()
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := DB.Create(&user).Error; err != nil {
		return nil
	}

	return &user
}

func GetToken() string {
	user := CreateUser("token", "token@gmail.com", "token")

	reqBody := map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/auth/login")
	LoginController(c)

	body := rec.Body.String()

	var response struct {
		User models.UserLogin `json:"data"`
	}

	_ = json.Unmarshal([]byte(body), &response)

	return fmt.Sprintf("Bearer %s", response.User.Token)
}

func TruncateUsersTable() {
	DB := database.Connect()
	DB.Exec("TRUNCATE TABLE users")
}

func TestCreateUserController(t *testing.T) {
	database.InitialMigration()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "create user normal",
			path:       "/users",
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"name":     "test1",
		"email":    "test1@gmail.com",
		"password": "test1",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				User models.User `json:"user"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.User.Name, "test1")
			assert.Equal(t, response.User.Email, "test1@gmail.com")
			assert.Equal(t, response.User.Password, "test1")
		}
	}

	TruncateUsersTable()
}

func TestGetUsersController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()

	users := []models.User{
		{
			Name: "test1", Email: "test1@gmail.com", Password: "test1",
		},
		{
			Name: "test2", Email: "test2@gmail.com", Password: "test2",
		},
	}

	for _, user := range users {
		CreateUser(user.Name, user.Email, user.Password)
	}

	testCases := []struct {
		name       string
		path       string
		expectCode int
		dataSize   int
	}{
		{
			name:       "get all users normal",
			path:       "/restricted/users",
			expectCode: http.StatusOK,
			dataSize:   len(users) + 1, // + 1 ketambahan user buat login
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/restricted/users", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				Users []models.User `json:"users"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testCase.dataSize, len(response.Users))
		}
	}

	TruncateUsersTable()
}

func TestGetUserController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()
	user := CreateUser("test1", "test1@gmail.com", "test1")

	testCases := []struct {
		name       string
		path       string
		user_id    string
		expectCode int
	}{
		{
			name:       "get user normal",
			path:       "/restricted/users",
			user_id:    strconv.Itoa(int(user.ID)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/restricted/users", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.user_id)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				User models.User `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.User.Name, "test1")
			assert.Equal(t, response.User.Email, "test1@gmail.com")
			assert.Equal(t, response.User.Password, "test1")
		}
	}

	TruncateUsersTable()
}

func TestUpdateUserController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()
	user := CreateUser("test1", "test1@gmail.com", "test1")

	testCases := []struct {
		name       string
		path       string
		user_id    string
		expectCode int
	}{
		{
			name:       "update user normal",
			path:       "/restricted/users",
			user_id:    strconv.Itoa(int(user.ID)),
			expectCode: http.StatusOK,
		},
	}

	reqBody := map[string]interface{}{
		"name":     "updated_name",
		"email":    "updated_email@gmail.com",
		"password": "updated_password",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/restricted/users", bytes.NewReader(reqBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.user_id)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				User models.User `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.User.Name, "updated_name")
			assert.Equal(t, response.User.Email, "updated_email@gmail.com")
			assert.Equal(t, response.User.Password, "updated_password")
		}
	}

	TruncateUsersTable()
}

func TestDeleteUserController(t *testing.T) {
	database.InitialMigration()

	token := GetToken()
	user := CreateUser("test1", "test1@gmail.com", "test1")

	testCases := []struct {
		name       string
		path       string
		user_id    string
		expectCode int
	}{
		{
			name:       "delete user normal",
			path:       "/restricted/users",
			user_id:    strconv.Itoa(int(user.ID)),
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/restricted/users", nil)
	req.Header.Set("Authorization", token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.user_id)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response struct {
				User interface{} `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Nil(t, response.User)
		}
	}

	TruncateUsersTable()
}
