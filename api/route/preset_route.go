package route

import (
	"github.com/d1nnn/api/controller"
	"github.com/d1nnn/api/middleware"
	"github.com/d1nnn/repository"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewPresetRoute(e *echo.Group, db *gorm.DB) {
	pu := usecase.NewPresetUsecase(repository.NewPostgresPresetRepository(db))
	pc := controller.NewPresetController(pu)

	presetRoute := e.Group("/presets")
	presetRoute.Use(middleware.ClerkJwtMiddleware(), middleware.JwtAuthMiddleware)

	presetRoute.GET("/user/:userid", pc.GetAllFromUser)
	presetRoute.POST("", pc.CreatePreset)
	presetRoute.POST("/delete", pc.DeletePreset)
}
