package shared_infrastructure

import (
	"context"
	"errors"
	"fmt"
	"time"

	shared_domain "github.com/fedeegmz/auth-service/internal/shared/domain"
	"github.com/jmoiron/sqlx"
)

var ErrConnectingDatabase = errors.New("database cannot be connected")

type Database struct {
	*sqlx.DB
}

func ConnectDatabase(cfg shared_domain.DatabaseSettings) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, ErrConnectingDatabase
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		if err := db.Close(); err != nil {
			fmt.Printf("Error closing database during connection failure: %v", err)
		}
		return nil, ErrConnectingDatabase
	}

	return &Database{db}, nil
}
