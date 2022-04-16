package users

import (
	"belajar-go-echo/domains/users"

	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUserSQLiteRepository(db *gorm.DB) users.UserRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(user users.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repository) GetUsers() []users.User {
	users := []users.User{}
	r.DB.Find(&users)

	return users
}

func (r *repository) GetUser(id int) (users.User, error) {
	user := users.User{}
	res := r.DB.Where("id = ?", id).Find(&user)

	if res.RowsAffected < 1 {
		return user, res.Error
	}

	return user, nil
}
