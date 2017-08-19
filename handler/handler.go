package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func RenderRoot(c echo.Context) error {
	return view.Slim(c, http.StatusOK, "something.slim", map[string]interface{}{})
}
