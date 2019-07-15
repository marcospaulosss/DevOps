package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
)

func BindAndValidate(c echo.Context, item interface{}) error {
	log.Info("Validando a request...")
	if err := c.Bind(item); err != nil {
		log.Error("Falhou. Problemas no binding:", err)
		c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidJSON})
		return err
	}
	if err := c.Validate(item); err != nil {
		log.Error("Falhou. Request invalido:", err)
		msg := fmt.Sprintf("%s %s", ErrInvalidRequest, err.Error())
		c.JSON(http.StatusBadRequest, structs.Response{Err: msg})
		return err
	}

	return nil
}

func ValidateID(c echo.Context) (uint64, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		log.Error("Falhou. ID nao numerico ou menor que 1.", err)
		c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidID})
		return 0, errors.New(ErrInvalidID)
	}
	return uint64(id), err
}

func ValidateUUID(c echo.Context) (string, error) {
	id := c.Param("id")
	if len(id) < 36 {
		log.Error("Falhou. ID nao eh UUID.")
		c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidUUID})
		return "", errors.New("Invalid UUID")
	}
	return id, nil
}

func ValidateEmpty(c echo.Context, item string) error {
	if item == "" || item == " " {
		log.Error("Falhou. item esta em branco.")
		c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidRequest})
		return errors.New("Invalid Request")
	}
	return nil
}
