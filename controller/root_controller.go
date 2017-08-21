package controller

import (
	"net/http"

	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/view"
	"github.com/labstack/echo"
)

func RenderRoot(c echo.Context) error {
	return view.Slim(c, http.StatusOK, "root.slim", map[string]interface{}{})
}
