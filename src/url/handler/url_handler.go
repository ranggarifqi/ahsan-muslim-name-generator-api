package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ranggarifqi/qashir-api/helper"
	"github.com/ranggarifqi/qashir-api/src/response"
	"github.com/ranggarifqi/qashir-api/src/url"
)

type urlHandler struct {
	urlUsecase url.IUrlUsecase
}

func NewUrlHandler(g *echo.Group, urlUsecase url.IUrlUsecase) {
	handler := &urlHandler{
		urlUsecase,
	}

	g.GET("/url/:id", handler.FindById)
	g.POST("/url", handler.Create)
}

func (h *urlHandler) FindById(c echo.Context) error {
	id := c.Param("id")
	res, err := h.urlUsecase.FindById(id)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "Data fetched successfully",
		Data:       res,
	})
}

func (h *urlHandler) Create(c echo.Context) error {
	urlDto := new(url.CreateUrlDto)
	err := c.Bind(urlDto)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusInternalServerError)
	}

	if err = c.Validate(urlDto); err != nil {
		return helper.HandleHttpError(c, err, http.StatusBadRequest)
	}

	result, err := h.urlUsecase.Create(urlDto)
	if err != nil {
		return helper.HandleHttpError(c, err, http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, response.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "Data created successfully",
		Data:       *result,
	})
}
