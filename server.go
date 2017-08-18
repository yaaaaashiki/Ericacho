package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yaaaaashiki/Ericacho/db"
	"github.com/yaaaaashiki/Ericacho/handler"
	"github.com/yaaaaashiki/Ericacho/model"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "something.html", map[string]interface{}{})
	}).Name = "foobar"

	db := db.GormConnect()
	user := &model.User{}
	db.First(user)
	log.Printf("----------------------------------------")
	log.Println(user.Name)
	log.Printf("----------------------------------------")

	// Route
	e.GET("/hello", handler.MainPage())

	e.Logger.Fatal(e.Start(":8080"))
}
