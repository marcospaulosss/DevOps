package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type ShelfHandler struct {
	shelfService interfaces.Service
}

func NewShelfHandler(serviceContainer services.Container) *ShelfHandler {
	return &ShelfHandler{
		shelfService: serviceContainer.Shelf,
	}
}

func (this *ShelfHandler) Create(c echo.Context) error {
	var item structs.Shelf
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	saved, err := this.shelfService.Create(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate})
	}
	return c.JSON(http.StatusCreated, structs.Response{Data: saved})
}

func (this *ShelfHandler) Update(c echo.Context) error {
	var item structs.Shelf
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item.ID = id
	saved, err := this.shelfService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: saved})
}

func (this *ShelfHandler) ReadOne(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Shelf{ID: id}
	result, err := this.shelfService.ReadOne(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadOne})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: result})
}

func (this *ShelfHandler) FetchShelvesOnlyWithPublishedAlbums(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	search.Extra = `{"album__is_published": "true"}`
	result, err := this.shelfService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	var filtered []structs.Shelf
	for _, i := range result.Items.([]structs.Shelf) {
		if len(i.Albums) > 0 {
			filtered = append(filtered, i)
		}
	}
	resp.Data = filtered
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}

func (this *ShelfHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	result, err := this.shelfService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}

func (this *ShelfHandler) Delete(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Shelf{ID: id}
	_, err := this.shelfService.Delete(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrDelete})
	}
	return c.NoContent(http.StatusNoContent)
}
