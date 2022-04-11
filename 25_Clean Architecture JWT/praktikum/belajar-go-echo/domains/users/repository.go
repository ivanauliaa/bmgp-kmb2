package users

type UserRepository interface {
	CreateUser(user User) error
	GetUsers() []User
	GetUser(id int) (User, error)
}
