package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
