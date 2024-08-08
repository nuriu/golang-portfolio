package services

import (
	"errors"
	"task-manager/internal/domain/user"
)

type UserSevice struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) user.UserService {
	return &UserSevice{userRepository}
}

// CreateUser implements user.UserService.
func (u *UserSevice) CreateUser(email string, password string) (*user.User, error) {
	existingUser, _ := u.userRepository.Get(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	generatedUser, err := user.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	createdUser, err := u.userRepository.Create(generatedUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// GetUser implements user.UserService.
func (u *UserSevice) GetUser(email string) (*user.User, error) {
	userDetail, err := u.userRepository.Get(email)
	if err != nil {
		return nil, err
	}

	return userDetail, nil
}
