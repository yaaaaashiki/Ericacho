package interceptor

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "hoge" && password == "password" {
			return true, nil
		}
		return false, nil
	})
}
