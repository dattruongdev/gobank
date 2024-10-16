package main

import (
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/d1nnn/api/route"
	bootstrap "github.com/d1nnn/boostrap"
	"github.com/labstack/echo/v4"
)

func main() {
	app := bootstrap.NewApp()
	env := app.Env
	db := app.Db

	e := echo.New()

	clerk.SetKey(env.Clerk_Secret)

	route.SetupRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
