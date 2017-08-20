package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/suzuken/wiki/sessions"
	"github.com/yaaaaashiki/Ericacho/crypto"
	"github.com/yaaaaashiki/Ericacho/mail"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

type User struct {
	DB *gorm.DB
}

func (u *User) CreateUser(c echo.Context) error {
	user := &model.User{}

	user.Name = c.FormValue("username")
	user.Email = c.FormValue("email")
	user.Salt = crypto.Salt(100)
	user.Salted = crypto.Stretch(c.FormValue("password"), user.Salt)
	//user.Created, user.Updated = time.Now(), time.Now()
	user.Create(u.DB)

	u.NewUser(c)
	return nil
}

func (u *User) NewUser(c echo.Context) error {
	var user model.User
	users, err := user.FetchAllUsers(u.DB)
	if err != nil {
		log.Fatal(err)
	}

	usersData := map[string]interface{}{
		"names": users,
	}

	return view.Slim(c, http.StatusOK, "users/new.slim", usersData)
}

func (u *User) Index(c echo.Context) error {
	r := c.Request()
	sess, _ := sessions.Get(r, "user")
	fmt.Println(sess.Values["name"])

	emptyData := map[string]interface{}{
		"user": sess.Values["name"],
	}

	if addr, isString := sess.Values["email"].(string); isString {
		go mail.InfiniteMail(addr)
	} else {
		errors.New("This input is not email address")
	}

	return view.Slim(c, http.StatusOK, "users/index.slim", emptyData)
}
