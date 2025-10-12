package application

import (
	shared_domain "github.com/fedeegmz/auth-service/internal/shared/domain"
	"github.com/fedeegmz/auth-service/internal/users/application/dto"
	"github.com/fedeegmz/auth-service/internal/users/domain"
)

type SaveOneUseCase struct {
	Repository domain.UserRepository
}

func (uc *SaveOneUseCase) Execute(user dto.UserSaveDto) error {
	id, err := shared_domain.NewId()
	if err != nil {
		return err
	}

	domainUser := dto.UserSaveDtoToDomain(id, user)
	return uc.Repository.SaveOne(domainUser)
}
