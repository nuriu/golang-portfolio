package user_test

import (
	"task-manager/internal/domain/user"
	"testing"
)

func TestUser(t *testing.T) {
	email := "test@test.com"
	pass := "test"

	t.Run("NewUser should not accept empty email or empty password", func(t *testing.T) {
		u, err := user.NewUser("", pass)

		if err == nil {
			t.Error("NewUser should return error when received an empty email")
		}

		if u != nil {
			t.Error("NewUser should not return user data when received an empty email")
		}

		if err != user.ErrorUserEmailEmpty {
			t.Errorf("NewUser should return %s when received an empty email", user.ErrorUserEmailEmpty.Error())
		}

		u, err = user.NewUser(email, "")

		if err == nil {
			t.Error("NewUser should return error when received an empty password")
		}

		if u != nil {
			t.Error("NewUser should not return user data when received an empty password")
		}

		if err != user.ErrorUserPasswordEmpty {
			t.Errorf("NewUser should return %s when received an empty password", user.ErrorUserPasswordEmpty.Error())
		}
	})

	t.Run("CheckPassword should be able to check passwords correctly", func(t *testing.T) {
		u, err := user.NewUser(email, pass)
		if err != nil {
			t.Error("NewUser should not return any error when received a correct email and password")
		}

		if !u.CheckPassword(pass) {
			t.Errorf("CheckPassword should returned: %v when expected: %v", u.CheckPassword(pass), true)
		}

		if u.CheckPassword(email) {
			t.Errorf("CheckPassword should returned: %v when expected: %v", u.CheckPassword(email), false)
		}
	})
}
