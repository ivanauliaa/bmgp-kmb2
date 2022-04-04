package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple/database"
	"simple/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLoginController(t *testing.T) {
	database.InitialMigration()

	testCases := []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "login normal",
			path:       "/auth/login",
			expectCode: http.StatusOK,
		},
	}

	user := CreateUser("test1", "test1@gmail.com", "test1")

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

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, LoginController(c)) {
			body := rec.Body.String()

			var response struct {
				User models.UserLogin `json:"data"`
			}

			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, response.User.Name, "test1")
			assert.Equal(t, response.User.Email, "test1@gmail.com")
			assert.NotNil(t, response.User.Token)
		}
	}

	TruncateUsersTable()
}
