package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type TrackHandler struct {
	trackService interfaces.Service
}

func NewTrackHandler(serviceContainer services.Container) *TrackHandler {
	return &TrackHandler{
		trackService: serviceContainer.Track,
	}
}

func (this *TrackHandler) Create(c echo.Context) error {
	var item structs.Track
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}

	saved, err := this.trackService.Create(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate})
	}
	return c.JSON(http.StatusCreated, structs.Response{Data: saved})
}

func (this *TrackHandler) Update(c echo.Context) error {
	var item structs.Track
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item.ID = id
	saved, err := this.trackService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: saved})
}

func (this *TrackHandler) ReadOne(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Track{ID: id}
	result, err := this.trackService.ReadOne(item)
	found := result.(structs.Track)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadOne})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: &found})
}

func (this *TrackHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	result, err := this.trackService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}

func (this *TrackHandler) Delete(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Track{ID: id}
	_, err := this.trackService.Delete(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrDelete})
	}
	return c.NoContent(http.StatusNoContent)
}
