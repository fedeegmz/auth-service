package infrastructure

import (
	"errors"

	shared_infrastructure "github.com/fedeegmz/auth-service/internal/shared/infrastructure"
	"github.com/fedeegmz/auth-service/internal/users/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUsernameExists = errors.New("username exists")
)

type UserRepositoryImpl struct {
	db *shared_infrastructure.Database
}

func (r *UserRepositoryImpl) GetAll() []domain.User {
	return []domain.User{}
}

func (r *UserRepositoryImpl) GetOne(userId string) (domain.User, error) {
	return domain.User{}, ErrUserNotFound
}

func (r *UserRepositoryImpl) SaveOne(user domain.User) error {
	return nil
}

func NewUserRepositoryImpl(db *shared_infrastructure.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
