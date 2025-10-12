package shared_domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrCreatingId = errors.New("user id cannot be created")

func NewId() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", ErrCreatingId
	}

	return id.String(), nil
}
