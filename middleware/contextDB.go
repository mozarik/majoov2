package middleware

import (
	"github.com/labstack/echo/v4"
	postgres "github.com/mozarik/majoov2/internal/db"
)

// ContextDB : pass db
// e.Use(middlewares.ContextDB(db))
func ContextDB(db *postgres.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
