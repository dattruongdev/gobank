package middleware

import (
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		claims, ok := clerk.SessionClaimsFromContext(c.Request().Context())
		if !ok {
			return c.String(http.StatusUnauthorized, "You are unauthenticated")
		}

		c.Set("claims", claims)

		err := next(c)

		return err
	}
}

func ClerkJwtMiddleware() echo.MiddlewareFunc {
	return echo.WrapMiddleware(clerkhttp.WithHeaderAuthorization())
}