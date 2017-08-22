package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

/*
func RenderRoot(c echo.Context) error {
	return view.Slim(c, http.StatusOK, "root.slim", map[string]interface{}{})
}
*/

func Root(c echo.Context) error {
	return view.Ace(c, http.StatusOK, "base", "inner", map[string]interface{}{})
}
