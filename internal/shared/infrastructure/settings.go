package shared_infrastructure

import (
	"errors"
	"os"

	shared_domain "github.com/fedeegmz/auth-service/internal/shared/domain"
	"github.com/joho/godotenv"
)

var ErrLoadingSettings = errors.New("settings cannot be loaded")

func LoadSettings() (shared_domain.Settings, error) {
	err := godotenv.Load()
	if err != nil {
		return shared_domain.Settings{}, ErrLoadingSettings
	}

	settings := shared_domain.Settings{
		Port: os.Getenv("PORT"),

		Database: shared_domain.DatabaseSettings{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	return settings, nil
}
