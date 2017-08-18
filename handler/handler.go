package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

/*
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}
*/
