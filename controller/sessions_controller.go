package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/suzuken/wiki/sessions"
	"github.com/yaaaaashiki/Ericacho/crypto"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

func (u *User) NewSession(c echo.Context) error {
	return view.Ace(c, http.StatusOK, "sessions/new", "", map[string]interface{}{})
}

func (u *User) CreateSession(c echo.Context) error {
	db := db.GormConnect()
	var isLoginUser model.User

	db.Where("email = ?", c.FormValue("email")).Find(&isLoginUser)
	fmt.Println(isLoginUser)

	hashPassword := crypto.Stretch(c.FormValue("password"), isLoginUser.Salt)
	if isLoginUser.Salted != hashPassword {
		u.NewSession(c)
		return errors.New("Failed login, confirm email and password again")
	}

	r := c.Request()
	w := c.Response().Writer

	sess, _ := sessions.Get(r, "user")
	sess.Values["id"] = isLoginUser.Id
	sess.Values["email"] = isLoginUser.Email
	sess.Values["name"] = isLoginUser.Name
	if err := sessions.Save(r, w, sess); err != nil {
		log.Printf("session can't save: %s", err)
		return err
	}

	u.Index(c)

	return nil
}
