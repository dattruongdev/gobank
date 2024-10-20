package main

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/d1nnn/api/route"
	bootstrap "github.com/d1nnn/boostrap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := bootstrap.NewApp()
	env := app.Env
	db := app.Db

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
        AllowHeaders: []string{"Content-Type", "Authorization"},
    }))

	clerk.SetKey(env.Clerk_Secret)
	

	route.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
