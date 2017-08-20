package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

/*
func (u *User) Login(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return view.Default(w, r, http.StatusOK, "login.tmpl", nil)
	case "POST":
		return u.login(w, r)
	default:
		return &httputil.HTTPError{Status: http.StatusMethodNotAllowed}
	}
}
*/
func RenderSessionNew(c echo.Context) error {

	db := db.GormConnect()
	user := &model.User{}
	db.First(user)
	log.Println(user.Name)

	hogehoge := []string{}
	hogehoge = append(hogehoge, user.Name)
	huge := map[string]interface{}{
		"names": hogehoge,
	}

	defer db.Close()

	return view.Slim(c, http.StatusOK, "sessions/new.slim", huge)
}

func SessionCreate(c echo.Context) error {

	db := db.GormConnect()
	user := &model.User{}
	db.First(user)
	log.Println(user.Name)

	hogehoge := []string{}
	hogehoge = append(hogehoge, user.Name)
	huge := map[string]interface{}{
		"names": hogehoge,
	}

	defer db.Close()

	return view.Slim(c, http.StatusOK, "sessions/new.slim", huge)
}
