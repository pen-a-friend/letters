package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Health GET endpoint
func Health(c echo.Context) error {
	respBody := map[string]string{
		"status": "healthy",
	}
	return c.JSON(http.StatusOK, respBody)
}
