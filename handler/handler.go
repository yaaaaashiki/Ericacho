package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/view"
)

func RenderRoot(c echo.Context) error {
	hogehoge := []string{"foo", "bar", "baz", "hogehoge", "kjfdjkj"}
	huge := map[string]interface{}{
		"foo": hogehoge,
	}

	return view.Slim(c, http.StatusOK, "something.slim", huge)
}
