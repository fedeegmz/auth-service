package main

import (
	"net/http"

	"github.com/fedeegmz/auth-service/internal/users/application"
	"github.com/fedeegmz/auth-service/internal/users/domain"
	"github.com/fedeegmz/auth-service/internal/users/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userRepository := infrastructure.NewUserRepositoryImpl()

	e.GET("/users", func(c echo.Context) error {
		return handleGetUsers(c, userRepository)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		return handleGetUser(c, userRepository)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func handleGetUsers(c echo.Context, repository domain.UserRepository) error {
	useCase := application.GetAllUseCase{Repository: repository}
	return c.JSON(http.StatusOK, useCase.Execute())
}

func handleGetUser(c echo.Context, repository domain.UserRepository) error {
	userId := c.Param("id")
	useCase := application.GetOneUseCase{Repository: repository}

	user, err := useCase.Execute(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, user)
}
