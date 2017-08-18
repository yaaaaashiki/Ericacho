package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderRoot() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "something.html", map[string]interface{}{})
	}
}
