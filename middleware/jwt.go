package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func UseJWTAuth() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	return echoMiddleware.JWT([]byte(secret))
}
