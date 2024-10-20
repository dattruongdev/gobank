package controller

import (
	"net/http"

	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

func NewUserController(uc *usecase.UserUsecase) *UserController {
	return &UserController {
		userUsecase: uc,
	}
}

func (uc *UserController) GetAllUsersAsAdmin(c echo.Context) error {
	userid := c.Param("userid")
	users, err := uc.userUsecase.GetUsers(c.Request().Context(), userid)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response {
			Message: "Error fetching users from db",
			Status: 500,
		})
	}
	return c.JSON(http.StatusOK, users)
}