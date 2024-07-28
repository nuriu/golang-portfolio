package user

type UserRepository interface {
	Create(user *User) (*User, error)
	Get(email string) (*User, error)
}
