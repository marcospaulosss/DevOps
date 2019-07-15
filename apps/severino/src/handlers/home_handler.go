package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type HomeHandler struct {
	homeService interfaces.Service
}

func NewHomeHandler(serviceContainer services.Container) *HomeHandler {
	return &HomeHandler{
		homeService: serviceContainer.Home,
	}
}

func (this *HomeHandler) Create(c echo.Context) error {
	return nil
}

func (this *HomeHandler) Update(c echo.Context) error {
	return nil
}

func (this *HomeHandler) ReadOne(c echo.Context) error {
	return nil
}

func (this *HomeHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	result, err := this.homeService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}

func (this *HomeHandler) Delete(c echo.Context) error {
	return nil
}
