package dependencies

import (
	"errors"

	"github.com/fedeegmz/auth-service/internal/users/domain"
)

var ErrUserNotFound = errors.New("user not found")

type MockUserRepository struct {
	users []domain.User
}

func (r *MockUserRepository) GetAll() []domain.User {
	return r.users
}

func (r *MockUserRepository) GetOne(userId string) (domain.User, error) {
	for _, user := range r.users {
		if user.Id == userId {
			return user, nil
		}
	}
	return domain.User{}, ErrUserNotFound
}

func NewEmptyMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: []domain.User{
			{
				Id:             "1",
				Name:           "Wade",
				LastName:       "Wilson",
				Username:       "deadpool",
				HashedPassword: "plainpasswordfornow",
			},
			{
				Id:             "2",
				Name:           "Pedro",
				LastName:       "Parqueador",
				Username:       "peter",
				HashedPassword: "plainpasswordfornow",
			},
		},
	}
}
