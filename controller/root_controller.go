package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func Root(c echo.Context) error {
	return view.Ace(c, http.StatusOK, "base", "inner", map[string]interface{}{"Msg": "Hello Ace"})
}
