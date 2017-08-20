package main

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yaaaaashiki/Ericacho/controller"
)

type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "assets")
	e.GET("/", controller.RenderRoot)

	e.Static("/sessions", "assets")
	e.GET("/sessions/new", controller.RenderSessionNew)
	e.POST("/sessions", controller.SessionCreate)

	e.Logger.Fatal(e.Start(":8080"))
}
