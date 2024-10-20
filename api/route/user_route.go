package route

import (
	"github.com/d1nnn/api/controller"
	"github.com/d1nnn/api/middleware"
	"github.com/d1nnn/repository"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewUserRoute(e *echo.Group, db *gorm.DB) {
	uu := usecase.NewUserUsecase(repository.NewPostgresUserRepository(db))
	uc := controller.NewUserController(uu)

	userRoute := e.Group("/users")
	userRoute.Use(middleware.ClerkJwtMiddleware(), middleware.JwtAuthMiddleware)

	adminRoute := userRoute.Group("/admin")
	adminRoute.Use(middleware.ClerkJwtMiddleware(), middleware.JwtAuthMiddleware, middleware.WithAdminRole)
	adminRoute.GET("/:userid", uc.GetAllUsersAsAdmin)
}