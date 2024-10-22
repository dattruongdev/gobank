package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	signupUsecase *usecase.SignupUsecase
	userUsecase   *usecase.UserUsecase
}

func NewAuthController(uc *usecase.SignupUsecase, uu *usecase.UserUsecase) *AuthController {
	return &AuthController{
		signupUsecase: uc,
		userUsecase:   uu,
	}
}

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (ac *AuthController) SignUp(c echo.Context) error {
	jsonBody := make(map[string]interface{})

	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := jsonBody["UserID"].(string)

	fullName := jsonBody["FullName"].(string)
	email := jsonBody["Email"].(string)
	user, err := ac.userUsecase.GetByEmail(c.Request().Context(), email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Fetching user from db failed")
	}
	log.Println(user.Email, email)
	if user.Email == email {
		return c.String(http.StatusConflict, "User already exists")
	}

	request := usecase.SignupRequest{
		UserID:   userId,
		FullName: fullName,
		Email:    email,
	}

	err = ac.signupUsecase.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Creating user failed")
	}

	return c.JSON(http.StatusCreated, Response{
		Message: "User signed up",
		Status:  201,
	})
}
