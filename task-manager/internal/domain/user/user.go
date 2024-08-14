package user

import (
	"strings"
	"task-manager/internal/domain"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorUserEmailEmpty    = domain.NewDomainError(400, "email is empty")
	ErrorUserPasswordEmpty = domain.NewDomainError(400, "password is empty")
	ErrorUserAlreadyExists = domain.NewDomainError(400, "user already exists")
	ErrorUserNotFound      = domain.NewDomainError(404, "user not found")
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ID        uuid.UUID
	Email     string
	Password  string
}

func NewUser(email string, password string) (*User, error) {
	email = strings.Trim(email, " ")
	if len(email) < 1 {
		return nil, ErrorUserEmailEmpty
	}

	password = strings.Trim(password, " ")
	if len(password) < 1 {
		return nil, ErrorUserPasswordEmpty
	}

	return &User{
		ID:        uuid.Nil,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}, nil
}

func (user *User) CheckPassword(password string) bool {
	return user.Password == password
}
