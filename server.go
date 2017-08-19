package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/yaaaaashiki/Ericacho/handler"
)

type Template struct {
	templates *template.Template
}

/*
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
*/
type TemplateExecutor interface {
	ExecuteTemplate(w io.Writer, name string, data interface{}) error
}

func main() {
	e := echo.New()

	/*
		renderer := &Template{
			templates: template.Must(template.ParseGlob("templates/*.slim")),
		}
		e.Renderer = renderer
	*/
	e.GET("/", handler.RenderRoot)
	/*
			fs := make(map[string]slim.Func)
				hoge := []string{"世界", "日本"}

					fs := map[string]slim.Func {
						"add": func([]string{"世界", "日本"}){ return "a", nil},
						"sub": func([]string{"世界", "日本"}){ return "b", nil},
					}

					fmt.Println(hoge)
		commits := map[string]interface{}{
			"rsc": 3711,
			"r":   2138,
			"gri": 1908,
			"adg": 912,
		}
		view.Init(fs)
		view.Execute(os.Stdout, "/Users/Soichiro/go/src/github.com/yaaaaashiki/Ericacho/templates/something.slim", commits)
	*/

	//	e.GET("/", handler.RenderRoot())

	//log.Println(controller.GetFirstUser())
	e.Logger.Fatal(e.Start(":8080"))
}
