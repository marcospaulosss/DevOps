package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type AlbumHandler struct {
	albumService interfaces.AlbumService
}

func NewAlbumHandler(serviceContainer services.Container) *AlbumHandler {
	return &AlbumHandler{
		albumService: serviceContainer.Album,
	}
}

func (this *AlbumHandler) Create(c echo.Context) error {
	var item structs.Album
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	saved, err := this.albumService.Create(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate})
	}
	return c.JSON(http.StatusCreated, structs.Response{Data: saved})
}

func (this *AlbumHandler) Update(c echo.Context) error {
	var item structs.Album
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item.ID = id
	item.IsPublished = false
	saved, err := this.albumService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: saved})
}

func (this *AlbumHandler) ReadAllPublishedOnly(c echo.Context) error {
	log.Info("Buscando apenas os albums publicados.")
	search := structs.NewSearch(c.QueryParams())
	search.Extra = `{"is_published": "true"}`
	return this.readAll(c, search)
}

func (this *AlbumHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	return this.readAll(c, search)
}

func (this *AlbumHandler) readAll(c echo.Context, search structs.Search) error {
	result, err := this.albumService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	log.Info("Sucesso.")
	return c.JSON(http.StatusOK, resp)
}

func (this *AlbumHandler) ReadOne(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Album{ID: id}
	resp, err := this.albumService.ReadOne(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadOne})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: resp})
}

func (this *AlbumHandler) Delete(c echo.Context) error {
	id, res := ValidateID(c)
	if id == 0 {
		return res
	}
	item := structs.Album{ID: id}
	_, err := this.albumService.Delete(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrDelete})
	}
	return c.NoContent(http.StatusNoContent)
}

func (this *AlbumHandler) Publish(c echo.Context) error {
	id, err := ValidateID(c)
	if err != nil || id < 1 {
		return err
	}
	_, err = this.albumService.Publish(structs.Album{ID: id})
	if err != nil {
		return c.JSON(http.StatusNotFound, structs.Response{Err: ErrReadOne})
	}
	return c.NoContent(http.StatusNoContent)
}

func (this *AlbumHandler) Unpublish(c echo.Context) error {
	id, err := ValidateID(c)
	if err != nil || id < 1 {
		return err
	}
	_, err = this.albumService.Unpublish(structs.Album{ID: id})
	if err != nil {
		return c.JSON(http.StatusNotFound, structs.Response{Err: ErrReadOne})
	}
	return c.NoContent(http.StatusNoContent)
}
