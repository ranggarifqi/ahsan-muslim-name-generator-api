package testutil

import (
	"github.com/labstack/echo/v4"
	myValidator "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/validator"
)

func SetupServer() *echo.Echo {
	e := echo.New()
	e.Validator = myValidator.NewMyValidator()
	return e
}
