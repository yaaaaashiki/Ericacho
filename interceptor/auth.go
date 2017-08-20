package interceptor

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	Username = "hoge"
	Password = "hoge"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == Username && password == Password {
			return true, nil
		}
		return false, nil
	})
}
