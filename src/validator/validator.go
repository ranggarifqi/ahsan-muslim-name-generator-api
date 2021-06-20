package validator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type MyValidator struct {
	validator *validator.Validate
}

func NewMyValidator() *MyValidator {
	return &MyValidator{
		validator: validator.New(),
	}
}

func (v *MyValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
