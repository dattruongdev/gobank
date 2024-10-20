package route

import (
	"github.com/d1nnn/api/controller"
	"github.com/d1nnn/api/middleware"
	"github.com/d1nnn/repository"
	"github.com/d1nnn/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewTransactionRoute(group *echo.Group, db *gorm.DB) {

	uc := usecase.NewTransactionUsecase(repository.NewPostgresTransactionRepository(db))
	tc := controller.NewTransactionController(uc)

	route := group.Group("/transactions")
	
	route.Use(middleware.ClerkJwtMiddleware(), middleware.JwtAuthMiddleware)
	route.GET("/user/:userid", tc.GetByUserId)
	route.POST("/user/:userid", tc.CreateTransaction)
}