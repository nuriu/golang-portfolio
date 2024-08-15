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

	t.Run("should not accept empty email on GetUser", func(t *testing.T) {
		userDetail, err := userService.GetUser("")
		if userDetail != nil {
			t.Error("GetUser should not return user data when email is empty")
		}
		if err != user.ErrorUserEmailEmpty {
			t.Errorf("GetUser should return %s error when email is empty", user.ErrorUserEmailEmpty)
		}
	})

}
