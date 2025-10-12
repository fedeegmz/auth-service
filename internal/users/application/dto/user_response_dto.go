package dto

import "github.com/fedeegmz/auth-service/internal/users/domain"

type UserResponseDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
}

func FromDomain(u domain.User) UserResponseDto {
	return UserResponseDto{
		Id:       u.Id,
		Name:     u.Name,
		LastName: u.LastName,
		Username: u.Username,
	}
}
