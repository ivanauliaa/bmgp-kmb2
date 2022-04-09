package sqlite

import (
	"belajar-go-echo/domains/users"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	ConnectDB()
}

func ConnectDB() {
	DB, _ = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
}

func MigrateDB() error {
	return DB.AutoMigrate(
		users.User{},
	)
}
