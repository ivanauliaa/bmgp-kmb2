package auth

type AuthService interface {
	Login(credential Auth) (*string, int)
}
