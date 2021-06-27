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
