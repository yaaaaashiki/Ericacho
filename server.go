package main

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yaaaaashiki/Ericacho/controller"
	"github.com/yaaaaashiki/Ericacho/interceptor"
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
	e.GET("/sessions/new", controller.RenderSessionNew, interceptor.BasicAuth())

	e.Logger.Fatal(e.Start(":8080"))
}
