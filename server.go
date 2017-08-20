package main

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/handler"
)

type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()

	e.Static("/", "assets")
	e.GET("/", handler.RenderRoot)
	e.Static("/sessions", "assets")
	e.GET("/sessions/new", handler.RenderSessionNew)

	e.Logger.Fatal(e.Start(":8080"))
}
