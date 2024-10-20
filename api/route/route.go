package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/api/v1")
	
	NewAuthRoute(v1, db)
	NewTransactionRoute(v1, db)
	NewUserRoute(v1, db)
}