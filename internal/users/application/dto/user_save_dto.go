package dto

import "github.com/fedeegmz/auth-service/internal/users/domain"

type UserSaveDto struct {
	Name     string `json:"name"      form:"name"`
	LastName string `json:"last_name" form:"last_name"`
	Username string `json:"username"  form:"username"`
	Password string `json:"password"  form:"password"`
}

func UserSaveDtoToDomain(id string, u UserSaveDto) domain.User {
	return domain.User{
		Id:             id,
		Name:           u.Name,
		LastName:       u.LastName,
		Username:       u.Username,
		HashedPassword: u.Password, // TODO: Hashear el password antes de guardarlo
	}
}
