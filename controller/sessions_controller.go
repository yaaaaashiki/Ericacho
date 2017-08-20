package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

func (u *User) RenderSessionNew(c echo.Context) error {
	var user model.User
	users, err := user.FetchAllUsers(u.DB)
	if err != nil {
		log.Fatal(err)
	}
	/*
		db := db.GormConnect()
		user := &model.User{}
		db.First(user)
		log.Println(user.Name)
	*/
	huge := map[string]interface{}{
		"names": users,
	}

	return view.Slim(c, http.StatusOK, "users/new.slim", huge)
}
