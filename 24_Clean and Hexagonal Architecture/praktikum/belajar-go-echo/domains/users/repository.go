package users

type UserRepository interface {
	CreateUser(user User) error
	GetUsers() []User
}
