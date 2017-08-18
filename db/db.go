package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	DBMS = "mysql"
	USER = "root"
	PASS = ""
	//PROTOCOL = "tcp(##.###.##.###:3306)"
	DBNAME = "user"

	CONNECT = USER + ":" + PASS + "@" + "/" + erikachan
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
