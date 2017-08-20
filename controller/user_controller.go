package controller

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/crypto"
	"github.com/yaaaaashiki/Ericacho/model"
)

type User struct {
	DB *gorm.DB
}

func (u *User) CreateUser(c echo.Context) error {
	user := &model.User{}

	user.Name = "yashikihogehogeyashiki"
	user.Email = c.FormValue("email")
	user.Salt = crypto.Salt(100)
	user.Salted = crypto.Stretch(c.FormValue("password"), user.Salt)
	user.Created, user.Updated = time.Now(), time.Now()
	user.Create(u.DB)

	u.RenderSessionNew(c)
	return nil
}
