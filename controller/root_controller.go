package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
)

func GetFirstUser() *gorm.DB {
	user := &model.User{}
	db := db.GormConnect()
	firstUser := db.First(user)
	return firstUser
}
