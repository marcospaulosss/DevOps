package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
)

type ProductHandler struct {
	productService interfaces.Service
}

func NewProductHandler(serviceContainer services.Container) *ProductHandler {
	return &ProductHandler{
		productService: serviceContainer.Product,
	}
}

func (this *ProductHandler) Create(c echo.Context) error {
	var item structs.Product
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	saved, err := this.productService.Create(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate})
	}
	result := saved.(structs.Product)
	return c.JSON(http.StatusCreated, structs.Response{Data: map[string]string{"id": result.ID}})
}

func (this *ProductHandler) Update(c echo.Context) error {
	var item structs.Product
	if err := BindAndValidate(c, &item); err != nil {
		return err
	}
	id, res := ValidateUUID(c)
	if id == "" {
		return res
	}
	item.ID = id
	saved, err := this.productService.Update(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrUpdate})
	}
	result := saved.(structs.Product)
	return c.JSON(http.StatusOK, structs.Response{Data: map[string]string{"id": result.ID}})
}

func (this *ProductHandler) ReadOne(c echo.Context) error {
	id, res := ValidateUUID(c)
	if id == "" {
		return res
	}
	item := structs.Product{ID: id}
	result, err := this.productService.ReadOne(item)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadOne})
	}
	return c.JSON(http.StatusOK, structs.Response{Data: result})
}

func (this *ProductHandler) ReadAll(c echo.Context) error {
	search := structs.NewSearch(c.QueryParams())
	result, err := this.productService.ReadAll(search)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrReadAll})
	}
	resp := structs.Response{}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}
