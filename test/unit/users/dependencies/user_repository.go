package dependencies

import (
	"github.com/fedeegmz/auth-service/internal/users/domain"
	"github.com/fedeegmz/auth-service/internal/users/infrastructure"
)

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
	return domain.User{}, infrastructure.ErrUserNotFound
}

func (r *MockUserRepository) SaveOne(user domain.User) error {
	for _, u := range r.users {
		if u.Username == user.Username {
			return infrastructure.ErrUsernameExists
		}
	}
	r.users = append(r.users, user)
	return nil
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
