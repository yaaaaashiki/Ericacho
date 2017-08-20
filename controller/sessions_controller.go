package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/crypto"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

func (u *User) NewSession(c echo.Context) error {
	emptyData := map[string]interface{}{}

	return view.Slim(c, http.StatusOK, "sessions/new.slim", emptyData)
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

	u.Index(c)

	//user.Salted = crypto.Stretch(c.FormValue("password"), user.Salt)
	//user.Created, user.Updated = time.Now(), time.Now()

	//user.Create(u.DB)

	return nil
}
