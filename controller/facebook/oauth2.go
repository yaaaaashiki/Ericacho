package controllersFacebook

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/facebook"
)

func RenderFacebook(c echo.Context) error {
	config := facebook.GetConnect()

	url := config.AuthCodeURL("")
	RedirectHandler(url)

	return nil
}

func RedirectHandler(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, 301)
}
