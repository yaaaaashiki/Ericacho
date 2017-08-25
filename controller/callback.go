package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/facebook"
	"github.com/yaaaaashiki/Ericacho/view"

	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
)

func RenderFacebook(c echo.Context) error {
	log.Println("-----------------------------------------------")
	config := facebook.GetConnect()
	url := config.AuthCodeURL("")
	c.Redirect(301, url)
	return nil /*
		return view.Ace(c, http.StatusOK, "base", "inner", map[string]interface{}{"Msg": "Hello Ace"})
	*/
}

func Callback(c echo.Context) error {
	r := c.Request()
	config := facebook.GetConnect()

	tok, err := config.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		return errors.New("Cannnot converts an authorization code into a token.")
	}

	if tok.Valid() == false {
		return errors.New("vaild token")
	}

	client := config.Client(oauth2.NoContext, tok)
	session := &fb.Session{
		Version:    "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		return errors.New("Cannot returns facebook graph api call result")
	}

	userData := map[string]interface{}{
		"Id":    res["id"],
		"Name":  res["name"],
		"Email": res["email"],
	}

	return view.Ace(c, http.StatusMovedPermanently, "facebook/callback", "", userData)
}
