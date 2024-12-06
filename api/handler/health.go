package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "OK",
		"message": "API is running smoothly",
	})
}
