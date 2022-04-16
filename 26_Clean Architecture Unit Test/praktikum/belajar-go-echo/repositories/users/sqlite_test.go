package users_test

import (
	"belajar-go-echo/infrastructures/database"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetUsers(t *testing.T) {
	test := struct {
		name string
	}{
		name: "success",
	}

	t.Run(test.name, func(t *testing.T) {
		db_mock, mk, _ := sqlmock.New()
		database.DB, _ = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

		mk.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books`")).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "title", "author", "year"}).
					AddRow(1, "GRIT", "Mark Manson", "2018").
					AddRow(2, "Subtle of Not Giving a F", "Mark Manson", "2018"),
			)

		// TODO: Panggil get users

		assert.Equal(t, http.StatusOK, resBody.Status)
		assert.Equal(t, 2, len(resBody.Data))
	})
}
