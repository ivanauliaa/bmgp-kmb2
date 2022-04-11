package auth

type AuthRepository interface {
	Login(credential Auth) error
}
