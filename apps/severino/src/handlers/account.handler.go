package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/libs/authentication"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/errors"
	log "backend/libs/logger"
)

type AccountHandler struct {
	accountService interfaces.Service
}

func NewAccountHandler(serviceContainer services.Container) *AccountHandler {
	return &AccountHandler{
		accountService: serviceContainer.Account,
	}
}

func (this *AccountHandler) CreateAccount(c echo.Context) error {
	account := new(structs.Account)
	if err := BindAndValidate(c, account); err != nil {
		return err
	}

	if account.Email == "" && account.Phone == "" {
		log.Error(ErrInvalidRequest)
		return c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidRequest})
	}

	if c.Request().RequestURI == "/auth/login/generate" {
		account.Exists = true
	}

	saved, err := this.accountService.Create(account)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrCreate + ": " + err.Error()})
	}

	return c.JSON(http.StatusNoContent, structs.Response{Data: saved})
}

func (this *AccountHandler) ValidateAccount(c echo.Context) error {
	Account := new(structs.Account)
	if err := BindAndValidate(c, Account); err != nil {
		return err
	}

	if Account.Email == "" && Account.Phone == "" {
		log.Error(ErrInvalidRequest)
		return c.JSON(http.StatusBadRequest, structs.Response{Err: ErrInvalidRequest})
	}

	found, err := this.accountService.ReadOne(Account)
	if err != nil {
		return c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: ErrInvalidRequest})
	}

	return c.JSON(http.StatusNoContent, structs.Response{Data: found})
}

func (this *AccountHandler) LoginApp(c echo.Context) error {
	found, err := this.login(c)
	if err != nil {
		return err
	}

	result := found.(structs.Account)
	user := structs.User{ID: result.UserID}
	token, err := authentication.GenerateToken(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Err: ErrCreate + ": " + err.Error()})
	}

	data := map[string]interface{}{
		"id":    user.ID,
		"token": token,
	}

	return c.JSON(http.StatusOK, structs.Response{Data: data})
}

func (this *AccountHandler) LoginWeb(c echo.Context) error {
	found, err := this.login(c)
	if err != nil {
		return err
	}

	result := found.(structs.Account)
	user := structs.User{ID: result.UserID}
	token, err := authentication.GenerateToken(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Err: ErrCreate + ": " + err.Error()})
	}

	cookie := authentication.GenerateCookieWithToken(token)
	c.SetCookie(cookie)

	data := map[string]interface{}{
		"id": user.ID,
	}

	return c.JSON(http.StatusOK, structs.Response{Data: data})
}

func (this *AccountHandler) login(c echo.Context) (interface{}, error) {
	var account structs.Account
	if err := BindAndValidate(c, &account); err != nil {
		return nil, err
	}

	var value string
	value = account.Phone
	if account.Type == "email" {
		value = account.Email
	}

	if err := ValidateEmpty(c, value); err != nil {
		return nil, err
	}

	account.Exists = true

	found, err := this.accountService.ReadOne(&account)
	if err != nil {
		c.JSON(errors.StatusCodeFrom(err), structs.Response{Err: err.Error()})
		return nil, err
	}

	return found, err
}
