package routes

import (
	"net/http"

	shared_infrastructure "github.com/fedeegmz/auth-service/internal/shared/infrastructure"
	"github.com/fedeegmz/auth-service/internal/users/application"
	"github.com/fedeegmz/auth-service/internal/users/application/dto"
	"github.com/fedeegmz/auth-service/internal/users/domain"
	"github.com/fedeegmz/auth-service/internal/users/infrastructure"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	e *echo.Echo
}

func NewUserRoutes(e *echo.Echo) *UserRoutes {
	return &UserRoutes{e}
}

func (r *UserRoutes) Configure(db *shared_infrastructure.Database) {
	userRepository := infrastructure.NewUserRepositoryImpl(db)

	r.e.GET("/users", func(c echo.Context) error {
		return handleGetUsers(c, userRepository)
	})

	r.e.GET("/users/:id", func(c echo.Context) error {
		return handleGetUser(c, userRepository)
	})

	r.e.POST("/users", func(c echo.Context) error {
		return handleSaveUser(c, userRepository)
	})
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

func handleSaveUser(c echo.Context, repository domain.UserRepository) error {
	dto := dto.UserSaveDto{}
	if err := c.Bind(&dto); err != nil {
		return c.String(http.StatusBadRequest, "user data is required")
	}

	useCase := application.SaveOneUseCase{Repository: repository}
	if err := useCase.Execute(dto); err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, "user created successfully")
}
