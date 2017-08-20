package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func RenderRoot(c echo.Context) error {
	return view.Slim(c, http.StatusOK, "root.slim", map[string]interface{}{})
}
