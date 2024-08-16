package services_test

import (
	"task-manager/internal/app/services"
	"task-manager/internal/db/repositories"
	"task-manager/internal/domain/user"
	"testing"

	gormSqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserService(t *testing.T) {
	db, err := gorm.Open(gormSqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Error("failed to open sqlite db")
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	email := "test@test.com"
	pass := "test"

	t.Run("CreateUser should be able to create a new user", func(t *testing.T) {
		createdUser, err := userService.CreateUser(email, pass)

		if err != nil {
			if err == user.ErrorUserAlreadyExists {
				t.Error("CreateUser should not return ErrorUserAlreadyExists when creating new user")
			} else {
				t.Error(err)
			}
		}

		if createdUser == nil {
			t.Error("CreateUser should be able to create a new user")
		}
	})

	t.Run("CreateUser should return correct error when user exists", func(t *testing.T) {
		createdUser, err := userService.CreateUser(email, pass)

		if err == nil || createdUser != nil {
			t.Error("CreateUser should return when user exists")
		}

		if err != user.ErrorUserAlreadyExists {
			t.Error("CreateUser should return ErrorUserAlreadyExists when creation attempt of existing user")
		}
	})

	t.Run("GetUser should not accept an empty email", func(t *testing.T) {
		userDetail, err := userService.GetUser("")

		if userDetail != nil {
			t.Error("GetUser should not return user data when email is empty")
		}
		if err != user.ErrorUserEmailEmpty {
			t.Errorf("GetUser should return %s error when email is empty", user.ErrorUserEmailEmpty)
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
