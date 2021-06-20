package helper

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/src/response"
)

func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%v: %v\n", msg, err)
	}
}

// HandleHttpError ...
func HandleHttpError(c echo.Context, err error, code int) error {
	statusCode := code
	message := err.Error()

	if e, ok := err.(*echo.HTTPError); ok {
		statusCode = e.Code
		message = fmt.Sprint(e.Message)
	}

	return c.JSON(statusCode, response.ErrorResponse{StatusCode: statusCode, Message: message})
}
