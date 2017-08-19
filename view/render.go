package view

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/mattn/go-slim"
)

func Slim(c echo.Context, code int, name string, data map[string]interface{}) error {

	tmpl, err := slim.ParseFile(fmt.Sprintf("templates/%s", name))
	if err != nil {
		return err
	}

	if err = tmpl.Execute(c.Response().Writer, data); err != nil {
		return err
	}
	c.Response().WriteHeader(code)

	return nil
}
