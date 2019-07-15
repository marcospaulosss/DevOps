package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type SubjectHandler struct {
	subjectService interfaces.Service
}

func NewSubjectHandler(serviceContainer services.Container) *SubjectHandler {
	return &SubjectHandler{
		subjectService: serviceContainer.Subject,
	}
}

func (this *SubjectHandler) Create(c echo.Context) error {
	var item structs.Subject
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	saved, err := this.subjectService.Create(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate})
	}
	return c.JSON(http.StatusCreated, structs.Response{Data: saved})
}

func (this *SubjectHandler) Update(c echo.Context) error {
	var item structs.Subject
	var err error
	if err = BindAndValidate(c, &item); err != nil {
		return err
	}

	item.ID, err = ValidateID(c)
	if err != nil {
		return err
	}

	saved, err := this.subjectService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: saved})
}

func (this *SubjectHandler) ReadOne(c echo.Context) error {
	id, err := ValidateID(c)
	if err != nil {
		return err
	}
	item := structs.Subject{ID: id}
	result, err := this.subjectService.ReadOne(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadOne})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: result})
}

func (this *SubjectHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	result, err := this.subjectService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		resp.Err = ErrReadAll
		return c.JSON(errors.StatusCodeFrom(err), resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}

func (this *SubjectHandler) Delete(c echo.Context) error {
	id, err := ValidateID(c)
	if err != nil {
		return err
	}
	item := structs.Subject{ID: id}
	_, err = this.subjectService.Delete(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrDelete})
	}
	return c.JSON(http.StatusNoContent, nil)
}
