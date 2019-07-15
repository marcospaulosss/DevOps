package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"backend/apps/severino/src/interfaces"
	"backend/apps/severino/src/libs/authentication"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	log "backend/libs/logger"
)

type AuthHandler struct {
	userService    interfaces.Service
	accountService interfaces.Service
}

const cookieName string = "token"

type JwtClaims struct {
	jwt.StandardClaims
}

func NewAuthHandler(serviceContainer services.Container) *AuthHandler {
	return &AuthHandler{
		userService:    serviceContainer.User,
		accountService: serviceContainer.Account,
	}
}

func (this *AuthHandler) Logout(c echo.Context) error {
	cookie := authentication.GenerateCookieWithToken("")
	yesterday := time.Now().AddDate(0, -0, -1)
	cookie.Expires = yesterday
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "Bye!")
}

func (this *AuthHandler) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) == 2 {
			reqToken = splitToken[1]
		}

		_, errToken := authentication.ParseJwtToken(reqToken)

		cookie := this.getTokenFromCookie(c)
		_, errCookie := authentication.ParseJwtToken(cookie)

		if errToken != nil && errCookie != nil {
			return c.JSON(http.StatusUnauthorized, structs.Response{Err: ErrUnauthorized})
		}

		return next(c)
	}
}

func (this *AuthHandler) ReturnValidate(c echo.Context) error {
	return c.JSON(http.StatusOK, "Logado com Sucesso")
}

func (this *AuthHandler) GetAuthenticatedUser(c echo.Context) error {
	return c.String(http.StatusOK, "User ID 123")
}

func (this *AuthHandler) getTokenFromCookie(c echo.Context) string {
	cookie, err := c.Cookie(authentication.CookieName)
	if err != nil {
		log.Error("Cookie not found.", err)
		return ""
	}
	return cookie.Value
}
