package controllersFacebook

import (
	"errors"

	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/facebook"

	"fmt"

	fb "github.com/huandu/facebook"
	"golang.org/x/oauth2"
)

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Code  string `form:"code"`
	State int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		panic(err)
	}

	fmt.Print("hogehogew")
	config := facebook.GetConnect()

	tok, err := config.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}

	client := config.Client(oauth2.NoContext, tok)
	session := &fb.Session{
		Version:    "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		panic(err)
	}

	c.Data["ID"] = res["id"]
	c.Data["Name"] = res["name"]
	c.Data["Email"] = res["email"]
	c.TplName = "facebook/callback.slim"
}
