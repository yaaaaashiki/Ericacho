package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func RenderRoot() echo.HandlerFunc {
	return func(c echo.Context, w http.ResponseWriter, r *http.Request) error {
		return view.Default(w, r, http.StatusOK, "something.slim", map[string]interface{}{})
	}
}

/*
func RenderRoot() error {
	return func(c echo.Context, w http.ResponseWriter, r *http.Request) error {
		return view.Default(w, r, http.StatusOK, "something.slim", map[string]interface{}{}),
	}
}
*/
/*
func RenderRoot(c echo.Context) error {
	return view.Slim(c, http.StatusOK, "test.slim", map[string]interface{}{})
*/
