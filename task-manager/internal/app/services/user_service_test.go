package services_test

import (
	"errors"
	"task-manager/internal/app/services"
	"task-manager/internal/db/repositories"
	"task-manager/internal/domain/user"
	"testing"

	gormSqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserServiceTests() (user.UserService, error) {
	db, err := gorm.Open(gormSqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to open sqlite db")
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	return userService, nil
}

func TestCreateUser(t *testing.T) {
	userService, err := setupUserServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	email := "test@test.com"
	pass := "test"

	t.Run("CreateUser should not create an user with empty email", func(t *testing.T) {
		createdUser, err := userService.CreateUser("", pass)

		if err == nil {
			t.Error("CreateUser should return error when received an empty email")
		}

		if createdUser != nil {
			t.Error("CreateUser should not return user data when received an empty email")
		}

		if err != user.ErrorUserEmailEmpty {
			t.Errorf("CreateUser should return %s when received an empty email", user.ErrorUserEmailEmpty.Error())
		}
	})

	t.Run("CreateUser should not create an user with empty password", func(t *testing.T) {
		createdUser, err := userService.CreateUser(email, "")

		if err == nil {
			t.Error("CreateUser should return error when received an empty password")
		}

		if createdUser != nil {
			t.Error("CreateUser should not return user data when received an empty password")
		}

		if err != user.ErrorUserPasswordEmpty {
			t.Errorf("CreateUser should return %s when received an empty password", user.ErrorUserPasswordEmpty.Error())
		}
	})

	t.Run("CreateUser should be able to create a new user", func(t *testing.T) {
		createdUser, err := userService.CreateUser(email, pass)

		if err != nil {
			if err == user.ErrorUserAlreadyExists {
				t.Errorf("CreateUser should not return %s when creating new user", user.ErrorUserAlreadyExists.Error())
			} else {
				t.Error(err)
			}
		}

		if createdUser == nil {
			t.Error("CreateUser should return created user")
		}
	})

	t.Run("CreateUser should return ErrorUserAlreadyExists when user exists", func(t *testing.T) {
		createdUser, err := userService.CreateUser(email, pass)

		if err == nil || createdUser != nil {
			t.Error("CreateUser should return when user exists")
		}

		if err != user.ErrorUserAlreadyExists {
			t.Error("CreateUser should return ErrorUserAlreadyExists when creation attempt of existing user")
		}
	})

}

func TestGetUser(t *testing.T) {
	userService, err := setupUserServiceTests()
	if err != nil {
		t.Error(err.Error())
	}

	email := "test@test.com"
	pass := "test"
	_, _ = userService.CreateUser(email, pass)
	t.Run("GetUser should not accept an empty email", func(t *testing.T) {
		userDetail, err := userService.GetUser("")

		if userDetail != nil {
			t.Error("GetUser should not return user data when email is empty")
		}
		if err != user.ErrorUserEmailEmpty {
			t.Error("GetUser should return ErrorUserEmailEmpty when email is empty")
		}
	})

	t.Run("GetUser should return correct user data", func(t *testing.T) {
		userDetail, err := userService.GetUser(email)

		if err != nil {
			t.Error("GetUser should not return any error when user exists")
		}

		if userDetail != nil {
			if userDetail.Email != email || userDetail.Password != pass {
				t.Error("GetUser should return correct user data")
			}
		}
	})

	t.Run("GetUser should return correct error when user not exists", func(t *testing.T) {
		userDetail, err := userService.GetUser("abc" + email)

		if err == nil {
			t.Error("GetUser should return correct error when user not exists")
		}

		if userDetail != nil {
			t.Error("GetUser should not return any user data when user not exists")
		}

		if err != user.ErrorUserNotFound {
			t.Errorf("GetUser should return correct error when user not exists")
		}
	})
}
