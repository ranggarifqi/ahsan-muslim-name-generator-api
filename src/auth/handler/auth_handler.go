package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/auth"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/response"
)

type AuthHandler struct {
	authUsecase auth.IAuthUsecase
}

func NewAuthHandler(g *echo.Group, auc auth.IAuthUsecase) {
	handler := &AuthHandler{
		authUsecase: auc,
	}

	g.POST("/signin", handler.SignIn)
}

func (h *AuthHandler) SignIn(c echo.Context) error {
	payload := new(auth.SignInDTO)
	err := c.Bind(payload)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusInternalServerError)
	}

	if err = c.Validate(payload); err != nil {
		return helper.HandleHttpError(c, err, http.StatusBadRequest)
	}

	res, err := h.authUsecase.SignIn(payload)
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
