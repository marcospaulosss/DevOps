package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/libs/authentication"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
	"backend/libs/util"
)

type UserHandler struct {
	userService interfaces.Service
}

func NewUserHandler(serviceContainer services.Container) *UserHandler {
	return &UserHandler{
		userService: serviceContainer.User,
	}
}

func (this *UserHandler) CreateUser(c echo.Context) error {
	user := new(structs.User)
	if err := BindAndValidate(c, user); err != nil {
		return err
	}

	saved, err := this.userService.Create(*user)
	if err != nil {
		log.Error("Nao consegui cadastrar o usuario: ", err.Error())

		st := status.Convert(err)

		responseCode := http.StatusNotFound
		if st.Code() != codes.NotFound {
			responseCode = http.StatusInternalServerError
		}

		return c.JSON(responseCode, structs.Response{Err: err.Error()})
	}

	savedUser := saved.(structs.User)
	token, err := authentication.GenerateToken(&savedUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.Response{Err: err.Error()})
	}

	responseBody := map[string]interface{}{
		"id":    savedUser.ID,
		"token": token,
	}

	log.Info("Sucesso. Criei o usuario:", util.ToJsonString(saved))
	return c.JSON(http.StatusCreated, structs.Response{Data: responseBody})
}

func (this *UserHandler) ReadAllUsers(c echo.Context) error {
	params := c.QueryParams()
	search := structs.NewSearch(params)
	result, err := this.userService.ReadAll(search)
	resp := structs.Response{}
	if err != nil {
		log.Error("Falhou. Nao achou nada.", err)
		resp.Err = ErrReadAll
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Data = result.Items
	resp.Meta = result.Pagination
	return c.JSON(http.StatusOK, resp)
}
