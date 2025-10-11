package application_test

import (
	"testing"

	"github.com/fedeegmz/auth-service/internal/users/application"
	"github.com/fedeegmz/auth-service/test/unit/users/dependencies"
)

func TestGetAllUseCase(t *testing.T) {
	t.Run("should return all users", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewMockUserRepository()
		useCase := application.GetAllUseCase{Repository: repository}

		// Act
		users := useCase.Execute()

		// Assert
		if len(users) != 2 {
			t.Fatalf("expected 2 users, got %d", len(users))
		}

		firstUser := users[0]
		if firstUser.Id != "1" {
			t.Errorf("first user id: expected \"1\", got %q", firstUser.Id)
		}
		if firstUser.Name != "Wade" {
			t.Errorf("first user name: expected \"Wade\", got %q", firstUser.Name)
		}
		if firstUser.LastName != "Wilson" {
			t.Errorf("first user lastname: expected \"Wilson\", got %q", firstUser.LastName)
		}
		if firstUser.Username != "deadpool" {
			t.Errorf("first user username: expected \"deadpool\", got %q", firstUser.Username)
		}

		secondUser := users[1]
		if secondUser.Id != "2" {
			t.Errorf("second user id: expected \"2\", got %q", secondUser.Id)
		}
		if secondUser.Name != "Pedro" {
			t.Errorf("second user name: expected \"Pedro\", got %q", secondUser.Name)
		}
		if secondUser.LastName != "Parqueador" {
			t.Errorf("second user lastname: expected \"Parqueador\", got %q", secondUser.LastName)
		}
		if secondUser.Username != "peter" {
			t.Errorf("second user username: expected \"peter\", got %q", secondUser.Username)
		}
	})

	t.Run("should return empty array", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewEmptyMockUserRepository()
		useCase := application.GetAllUseCase{Repository: repository}

		// Act
		users := useCase.Execute()

		// Assert
		if len(users) != 0 {
			t.Fatalf("expected 0 users, got %d", len(users))
		}
	})
}
