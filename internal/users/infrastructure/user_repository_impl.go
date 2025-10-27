package infrastructure

import (
	"errors"

	shared_infrastructure "github.com/fedeegmz/auth-service/internal/shared/infrastructure"
	"github.com/fedeegmz/auth-service/internal/users/domain"
)

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUsernameExists = errors.New("username exists")
	ErrUserNotSaved   = errors.New("user cannot be saved")
)

type UserDbModel struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	LastName string `db:"last_name"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepositoryImpl struct {
	db *shared_infrastructure.Database
}

func (r *UserRepositoryImpl) GetAll() []domain.User {
	const query = "SELECT * FROM users"
	result := []domain.User{}
	users := []UserDbModel{}
	if err := r.db.Select(&users, query); err != nil {
		return result
	}

	for _, user := range users {
		result = append(result, domain.User{
			Id:             user.Id,
			Name:           user.Name,
			LastName:       user.LastName,
			Username:       user.Username,
			HashedPassword: user.Password,
		})
	}

	return result
}

func (r *UserRepositoryImpl) GetOne(userId string) (domain.User, error) {
	const query = "SELECT * FROM users u WHERE u.id = $1"
	user := UserDbModel{}
	err := r.db.Get(&user, query, userId)
	if err != nil {
		return domain.User{}, ErrUserNotFound
	}

	return domain.User{
		Id:             user.Id,
		Name:           user.Name,
		LastName:       user.LastName,
		Username:       user.Username,
		HashedPassword: user.Password,
	}, nil
}

func (r *UserRepositoryImpl) GetOneByUsername(username string) (domain.User, error) {
	const query = "SELECT * FROM users u WHERE u.username = $1"
	user := UserDbModel{}
	err := r.db.Get(&user, query, username)
	if err != nil {
		return domain.User{}, ErrUserNotFound
	}

	return domain.User{
		Id:             user.Id,
		Name:           user.Name,
		LastName:       user.LastName,
		Username:       user.Username,
		HashedPassword: user.Password,
	}, nil
}

func (r *UserRepositoryImpl) SaveOne(user domain.User) error {
	if _, err := r.GetOneByUsername(user.Username); err == nil {
		return ErrUsernameExists
	}

	const query = `
	INSERT INTO users (id, name, last_name, username, password)
	VALUES (:id, :name, :last_name, :username, :password)
	`
	data := UserDbModel{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
		Username: user.Username,
		Password: user.HashedPassword,
	}
	result, err := r.db.NamedExec(query, data)
	if err != nil {
		return ErrUserNotSaved
	}

	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return ErrUserNotSaved
	}

	return nil
}

func NewUserRepositoryImpl(db *shared_infrastructure.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
