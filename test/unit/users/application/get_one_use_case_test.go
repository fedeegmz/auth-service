package application_test

import (
	"errors"
	"testing"

	"github.com/fedeegmz/auth-service/internal/users/application"
	"github.com/fedeegmz/auth-service/test/unit/users/dependencies"
)

func TestGetOneUseCase(t *testing.T) {
	t.Run("should return one user if exist", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewMockUserRepository()
		useCase := application.GetOneUseCase{Repository: repository}

		// Act
		user, err := useCase.Execute("1")
		// Assert
		if err != nil {
			t.Fatalf("expected 1 user, got %d", err)
		}

		if user.Id != "1" {
			t.Errorf("first user id: expected \"1\", got %q", user.Id)
		}
		if user.Name != "Wade" {
			t.Errorf("first user name: expected \"Wade\", got %q", user.Name)
		}
		if user.LastName != "Wilson" {
			t.Errorf("first user lastname: expected \"Wilson\", got %q", user.LastName)
		}
		if user.Username != "deadpool" {
			t.Errorf("first user username: expected \"deadpool\", got %q", user.Username)
		}
	})

	t.Run("should return error if user not exist", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewMockUserRepository()
		useCase := application.GetOneUseCase{Repository: repository}

		// Act
		_, err := useCase.Execute("3")

		// Assert
		if !errors.Is(err, dependencies.ErrUserNotFound) {
			t.Errorf("user with id 3: expected %q, got %q", dependencies.ErrUserNotFound, err)
		}
	})
}
