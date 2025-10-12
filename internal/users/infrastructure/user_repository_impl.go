package infrastructure

import (
	"errors"

	"github.com/fedeegmz/auth-service/internal/users/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUsernameExists = errors.New("username exists")
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
	return domain.User{}, ErrUserNotFound
}

func (r *UserRepositoryImpl) SaveOne(user domain.User) error {
	for _, u := range r.users {
		if u.Username == user.Username {
			return ErrUsernameExists
		}
	}
	r.users = append(r.users, user)
	return nil
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
