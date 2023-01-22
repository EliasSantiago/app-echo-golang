package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	InvalidParameters = func(c echo.Context, err error, key string) error {
		return c.JSON(http.StatusBadRequest, "Invalid query params")
	}
)
