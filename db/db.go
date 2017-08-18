package db

import (
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
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
