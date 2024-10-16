package controller

import (
	"net/http"

	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	signupUsecase *usecase.SignupUsecase
}

func NewAuthController(uc *usecase.SignupUsecase) *AuthController {
	return &AuthController {
		signupUsecase: uc,
	}
}

func(ac *AuthController) SignUp(c echo.Context) error {
	var request usecase.SignupRequest
	err := c.Bind(&request); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	err = ac.signupUsecase.CreateUser(c.Request().Context(), request); if err != nil {
		return c.String(http.StatusInternalServerError, "Creating user failed")
	}

	return err
}