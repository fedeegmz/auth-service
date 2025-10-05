package application

import (
	"github.com/fedeegmz/auth-service/internal/users/application/dto"
	"github.com/fedeegmz/auth-service/internal/users/domain"
)

type GetOneUseCase struct {
	Repository domain.UserRepository
}

func (uc *GetOneUseCase) Execute(userId string) (dto.UserResponse, error) {
	user, err := uc.Repository.GetOne(userId)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.FromDomain(user), nil
}
