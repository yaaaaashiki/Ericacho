package main

import (
	"errors"
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	slim "github.com/mattn/go-slim"
	"github.com/yaaaaashiki/Ericacho/controller"
	"github.com/yaaaaashiki/Ericacho/handler"
	"github.com/yaaaaashiki/Ericacho/view"
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

func hoge(string) (string, error) {
	err := errors.New("hogehoegkj")
	return "yashiki", err
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.slim")),
	}
	e.Renderer = renderer

	fs := map[string]slim.Func{}
	view.Init(fs)

	e.GET("/", handler.RenderRoot())

	log.Println(controller.GetFirstUser())
	e.Logger.Fatal(e.Start(":8080"))
}
