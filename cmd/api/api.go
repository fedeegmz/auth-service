package main

import (
	"fmt"
	"log"

	"github.com/fedeegmz/auth-service/cmd/api/routes"
	shared_infrastructure "github.com/fedeegmz/auth-service/internal/shared/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	settings, err := shared_infrastructure.LoadSettings()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := shared_infrastructure.ConnectDatabase(settings.Database)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	e := echo.New()
	routes.NewUserRoutes(e).Configure(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", settings.Port)))
}
