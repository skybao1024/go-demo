package authservice

type ServiceInterface interface {
	Register(username, password, email string) error
	Login(username, password string) (string, error)
}
