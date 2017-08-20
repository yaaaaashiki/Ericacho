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

func (u *User) CreateSession(c echo.Context) error {
	fmt.Println("---------------------------------------------------------------------------------------------")

	user := &model.User{}

	user.Id = 1
	user.Salt = "hoge"
	user.Name = "yashikihogehogeyashiki"
	user.Email = c.FormValue("email")
	user.Salted = c.FormValue("password")
	user.Created, user.Updated = time.Now(), time.Now()

	fmt.Println(user)
	log.Println(user)
	user.Create(u.DB)
	/*
		hogehoge := []string{}
		hogehoge = append(hogehoge, user.Name)
		huge := map[string]interface{}{
			"names": hogehoge,
		}
	*/

	u.RenderSessionNew(c)
	return nil
}
