package route

import (
	"github.com/d1nnn/api/controller"
	"github.com/d1nnn/api/middleware"
	"github.com/d1nnn/repository"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewAuthRoute(group *echo.Group, db *gorm.DB) {
	repo := repository.NewPostgresUserRepository(db)
	uc := usecase.NewSignupUsecase(repo)
	uu := usecase.NewUserUsecase(repo)
	ac := controller.NewAuthController(uc, uu)

	authRoute := group.Group("/auth")
	authRoute.Use(middleware.ClerkJwtMiddleware(), middleware.JwtAuthMiddleware)
	authRoute.POST("/signup", ac.SignUp)
	
}
