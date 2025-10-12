package domain

type UserRepository interface {
	GetOne(userId string) (User, error)

	GetAll() []User

	SaveOne(user User) error
}
