package users

type UserService interface {
	CreateUserService(user User) error
	GetUsersService() []User
	GetUserService(id int) (User, error)
}
