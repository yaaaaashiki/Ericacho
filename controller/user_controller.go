package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/model"
)

type User struct {
	DB *gorm.DB
}

func (u *User) CreateUser(c echo.Context) error {
	user := &model.User{}

	user.Salt = "hoge"
	user.Name = "yashikihogehogeyashiki"
	user.Email = c.FormValue("email")
	user.Salted = c.FormValue("password")
	user.Created, user.Updated = time.Now(), time.Now()

	fmt.Println(user)
	log.Println(user)
	user.Create(u.DB)

	u.RenderSessionNew(c)
	return nil
}
