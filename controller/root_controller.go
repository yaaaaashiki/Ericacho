package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
	"github.com/yaaaaashiki/Ericacho/view"
)

func GetFirstUser() *gorm.DB {
	user := &model.User{}
	db := db.GormConnect()
	firstUser := db.First(user)
	return firstUser
}

func RenderRoot(c echo.Context) error {
	hogehoge := []string{}
	hogehoge = append(hogehoge, "hogehoge")
	huge := map[string]interface{}{
		"foo": hogehoge,
	}

	return view.Slim(c, http.StatusOK, "root.slim", huge)
}
