package users

import (
	"belajar-go-echo/domains/users"
	"belajar-go-echo/infrastructures/database/sqlite"

	"fmt"
)

func CreateUser(user users.User) error {
	res := sqlite.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAllUsers() []users.User {
	users := []users.User{}
	sqlite.DB.Find(&users)

	return users
}

func Login(loginRequest users.LoginRequest) *users.User {
	user := users.User{}

	err := sqlite.DB.Where("email = ? AND password = ?", loginRequest.Email, loginRequest.Password).Find(&user).Error
	if err != nil {
		return nil
	}

	return &user
}
