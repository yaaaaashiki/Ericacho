package db

import (
	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBMS   = "mysql"
	USER   = "root"
	PASS   = ""
	DBNAME = "ericacho"
)

func GormConnect() *gorm.DB {
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	DBMigrate(db)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Comment{})
}
