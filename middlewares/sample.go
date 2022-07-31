package middlewares

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func SampleMiddleware() (echo.MiddlewareFunc, error) {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			iWannaThrowError := false

			if iWannaThrowError {
				return echo.NewHTTPError(http.StatusUnauthorized, "message")
			}

			log.Println("You are going through sample middleware")
			return next(c)
		}
	}, nil
}
