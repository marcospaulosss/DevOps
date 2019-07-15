package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type PreferenceHandler struct {
	preferenceService interfaces.Service
}

func NewPreferenceHandler(serviceContainer services.Container) *PreferenceHandler {
	return &PreferenceHandler{serviceContainer.Preference}
}

func (this *PreferenceHandler) Update(c echo.Context) error {
	var item structs.Preference
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}

	item.Type = c.Param("type")
	saved, err := this.preferenceService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: saved})
}
