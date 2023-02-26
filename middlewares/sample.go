package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/logger"
	"net/http"
)

func SampleMiddleware() (echo.MiddlewareFunc, error) {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			iWannaThrowError := false

			if iWannaThrowError {
				return echo.NewHTTPError(http.StatusUnauthorized, "message")
			}

			logger.Error("You are going through sample middleware")
			return next(c)
		}
	}, nil
}
