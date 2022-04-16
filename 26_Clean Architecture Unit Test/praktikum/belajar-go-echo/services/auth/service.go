package auth

import (
	"belajar-go-echo/domains/auth"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type service struct {
	repository auth.AuthRepository
}

func NewAuthService(repo auth.AuthRepository) auth.AuthService {
	return &service{
		repository: repo,
	}
}

func (s *service) Login(credential auth.Auth) (*string, int) {
	err := s.repository.Login(credential)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError
	}

	return &jwt, -1
}
