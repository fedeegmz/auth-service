package infrastructure

import (
	"errors"

	"github.com/fedeegmz/auth-service/internal/users/domain"
)

type UserRepositoryImpl struct {
	users []domain.User
}

func (r *UserRepositoryImpl) GetAll() []domain.User {
	return r.users
}

func (r *UserRepositoryImpl) GetOne(userId string) (domain.User, error) {
	for _, user := range r.users {
		if user.Id == userId {
			return user, nil
		}
	}
	return domain.User{}, errors.New("user not found")
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{
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
