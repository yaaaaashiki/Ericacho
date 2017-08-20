package main

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yaaaaashiki/Ericacho/controller"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/interceptor"
)

type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()

	db := db.GormConnect()
	defer db.Close()

	user := &controller.User{DB: db}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "assets")
	e.GET("/", controller.RenderRoot)

	e.Static("/users", "assets")
	e.GET("/users/new", user.RenderSessionNew, interceptor.BasicAuth())
	e.POST("/users", user.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
