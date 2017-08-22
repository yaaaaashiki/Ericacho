package view

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
	"github.com/yosssi/ace"
)

func Error(c echo.Context, err error) error {
	log.Fatalln(err)

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"errors": []string{"Server Internal Error"},
	})
}

func Slim(c echo.Context, code int, name string, data map[string]interface{}) error {
	tmpl, err := slim.ParseFile(fmt.Sprintf("templates/%s", name))
	if err != nil {
		c.Error(err)
		return nil
	}

	if err := tmpl.Execute(c.Response().Writer, data); err != nil {
		return Error(c, err)
	}
	c.Response().WriteHeader(code)

	return nil
}

func Ace(c echo.Context, code int, base, inner string, data map[string]interface{}) error {

	if base != "" {
		base = fmt.Sprintf("./templates/%s", base)
	}
	if inner != "" {
		inner = fmt.Sprintf("./templates/%s", inner)
	}

	tmpl, err := ace.Load(base, inner, &ace.Options{
		DynamicReload: true,
		DelimLeft:     "<%",
		DelimRight:    "%>",
	})
	if err != nil {
		return Error(c, err)
	}

	if err := tmpl.Execute(c.Response().Writer, data); err != nil {
		return Error(c, err)
	}
	c.Response().WriteHeader(code)

	return nil
}
