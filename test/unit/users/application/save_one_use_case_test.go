package application_test

import (
	"testing"

	"github.com/fedeegmz/auth-service/internal/users/application"
	"github.com/fedeegmz/auth-service/internal/users/application/dto"
	"github.com/fedeegmz/auth-service/internal/users/infrastructure"
	"github.com/fedeegmz/auth-service/test/unit/users/dependencies"
)

func TestSaveOneUseCase(t *testing.T) {
	t.Run("should return nil if user is valid", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewEmptyMockUserRepository()
		useCase := application.SaveOneUseCase{Repository: repository}
		dto := dto.UserSaveDto{
			Name:     "Tony",
			LastName: "Stark",
			Username: "ironman",
			Password: "iamironman",
		}

		// Act
		// Assert
		if err := useCase.Execute(dto); err != nil {
			t.Errorf("user is valid, expected nil error")
		}
	})

	t.Run("should return error if username exists", func(t *testing.T) {
		// Arrange
		repository := dependencies.NewMockUserRepository()
		useCase := application.SaveOneUseCase{Repository: repository}
		dto := dto.UserSaveDto{
			Name:     "Peter",
			LastName: "Parker",
			Username: "peter",
			Password: "iamironman",
		}

		// Act
		// Assert
		if err := useCase.Execute(dto); err != infrastructure.ErrUsernameExists {
			t.Errorf("username: expedted \"username exists\", got %q", err)
		}
	})
}
