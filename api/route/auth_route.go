package route

import (
	"github.com/d1nnn/api/controller"
	"github.com/d1nnn/repository"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewAuthRoute(group *echo.Group, db *gorm.DB) {
	uc := usecase.NewSignupUsecase(repository.NewPostgresUserRepository(db))
	ac := controller.NewAuthController(uc)

	authRoute := group.Group("/auth")
	authRoute.POST("/signup", ac.SignUp)
}