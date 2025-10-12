package dto

import "github.com/fedeegmz/auth-service/internal/users/domain"

type UserSaveDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserSaveDtoToDomain(id string, u UserSaveDto) domain.User {
	return domain.User{
		Id:       id,
		Name:     u.Name,
		LastName: u.LastName,
		Username: u.Username,
	}
}
