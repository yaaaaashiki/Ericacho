package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/model"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	db := db.GormConnect()
	// 構造体のインスタンス化
	user := &model.User{}
	// IDの指定
	user.Id = 1
	// 指定したIDを元にレコードを１つ引っ張ってくる
	db.First(user)
	console.log(user)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
