package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/response"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
)

type authHandler struct {
	userUsecase user.IUserUsecase
}

func NewAuthHandler(g *echo.Group, uuc user.IUserUsecase) {
	handler := &authHandler{
		userUsecase: uuc,
	}

	g.POST("/signin", handler.SignIn)
}

func (h *authHandler) SignIn(c echo.Context) error {
	payload := new(user.SignInDTO)
	err := c.Bind(payload)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusInternalServerError)
	}

	if err = c.Validate(payload); err != nil {
		return helper.HandleHttpError(c, err, http.StatusBadRequest)
	}

	res, err := h.userUsecase.SignIn(payload)
	if err != nil {
		err = errors.New("Incorrect email or password!")
		return helper.HandleHttpError(c, err, http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "User signed in successfully",
		Data:       *res,
	})
}
