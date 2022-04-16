package users

import (
	"belajar-go-echo/domains/users"
)

type service struct {
	repository users.UserRepository
}

func NewUserService(repo users.UserRepository) users.UserService {
	return &service{
		repository: repo,
	}
}

func (s *service) CreateUserService(user users.User) error {
	return s.repository.CreateUser(user)
}

func (s *service) GetUsersService() []users.User {
	return s.repository.GetUsers()
}

func (s *service) GetUserService(id int) (users.User, error) {
	return s.repository.GetUser(id)
}
