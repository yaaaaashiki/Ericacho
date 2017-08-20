package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func RenderRoot(c echo.Context) error {
	emptyData := map[string]interface{}{}

	return view.Slim(c, http.StatusOK, "root.slim", emptyData)
}
