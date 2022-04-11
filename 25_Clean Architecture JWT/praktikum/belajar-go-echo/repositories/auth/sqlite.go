package auth

import (
	"belajar-go-echo/domains/auth"
	"belajar-go-echo/domains/users"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthSQLiteRepository(db *gorm.DB) auth.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Login(credential auth.Auth) error {
	res := r.DB.Where("email = ? AND password = ?", credential.Email, credential.Password).Find(&users.User{})

	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}
