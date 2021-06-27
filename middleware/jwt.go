package middleware

import (
	"os"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	authService "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth/services"
)

func UseJWTAuth() echo.MiddlewareFunc {
	secret := os.Getenv("JWT_SECRET")
	return echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
		SigningKey: []byte(secret),
		Claims:     &authService.JwtClaim{},
	})
}
