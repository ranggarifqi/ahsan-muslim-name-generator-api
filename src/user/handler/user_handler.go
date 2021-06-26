package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/helper"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/response"
	"github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"
)

type userHandler struct {
	uc user.IUserUsecase
}

func NewUserHandler(g *echo.Group, uc user.IUserUsecase) {
	handler := &userHandler{uc}

	g.GET("/users/:id", handler.FindById)
}

func (uh *userHandler) FindById(c echo.Context) error {
	id := c.Param("id")
	res, err := uh.uc.FindById(id)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, response.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "Data fetched successfully",
		Data:       *res,
	})
}

func (uh *userHandler) SignIn(c echo.Context) error {
	payload := new(user.SignInDTO)
	err := c.Bind(payload)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusInternalServerError)
	}

	if err = c.Validate(payload); err != nil {
		return helper.HandleHttpError(c, err, http.StatusBadRequest)
	}

	res, err := uh.uc.SignIn(payload)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "User signed in successfully",
		Data:       *res,
	})
}
