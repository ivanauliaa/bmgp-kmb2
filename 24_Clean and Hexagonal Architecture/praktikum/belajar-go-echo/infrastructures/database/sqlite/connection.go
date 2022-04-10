package sqlite

import (
	"belajar-go-echo/domains/users"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db connection error")
		return nil
	}

	DB.AutoMigrate(
		users.User{},
	)

	return DB
}
