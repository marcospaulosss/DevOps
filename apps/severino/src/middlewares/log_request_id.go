package middlewares

import (
	log "backend/libs/logger"

	"github.com/labstack/echo/v4"
)

func LogRequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			log.SetRequestID(requestID)
			method := c.Request().Method
			log.Info("BEGIN", method, c.Path(), c.ParamValues(), c.QueryString())
			// c.Response().After(func() {
			// log.Info(c.Path())
			// })
			return next(c)
		}
	}
}
