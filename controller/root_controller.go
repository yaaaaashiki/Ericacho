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
/*
func Handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("base", "inner", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
*/
func Root(c echo.Context) error {
	return view.Ace(c, http.StatusOK, "base", "inner", map[string]interface{}{"Msg": "Hello Ace"})
}
