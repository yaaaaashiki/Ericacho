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

	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	e.GET("/", handler.RenderRoot)
	e.GET("/sessions/new", handler.RenderSessionNew)
	//e.Static("/", "assets/stylesheets")
	e.Logger.Fatal(e.Start(":8080"))
}
