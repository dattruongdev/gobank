package middleware

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/labstack/echo/v4"
)


func WithAdminRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok :=c.Get("claims").(*clerk.SessionClaims)

		if !ok {
			return c.String(http.StatusBadRequest, "Invalid claims")
		}

		if !claims.HasRole("org:admin") {
			return c.String(http.StatusForbidden, "You are not the admin")
		}

		err := next(c)

		return err
	}
}