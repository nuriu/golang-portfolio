package user

type UserService interface {
	CreateUser(email string, password string) (*User, error)
	GetUser(email string) (*User, error)
}
