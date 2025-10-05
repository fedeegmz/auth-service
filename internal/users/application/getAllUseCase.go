package application

import (
	"github.com/fedeegmz/auth-service/internal/users/application/dto"
	"github.com/fedeegmz/auth-service/internal/users/domain"
)

type GetAllUseCase struct {
	Repository domain.UserRepository
}

func (uc *GetAllUseCase) Execute() []dto.UserResponse {
	users := uc.Repository.GetAll()

	if len(users) == 0 {
		return []dto.UserResponse{}
	}

	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = dto.FromDomain(user)
	}
	return responses
}
