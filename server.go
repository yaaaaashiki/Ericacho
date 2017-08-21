package main

import (
	"html/template"

	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/controller"
	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/db"
	"github.com/VG-Tech-Dojo/treasure2017/mid/yaaaaashiki/interceptor"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
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
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Static("/", "assets")
	e.GET("/", controller.RenderRoot)

	e.Static("/users", "assets")
	e.GET("/users/new", user.NewUser, interceptor.BasicAuth())
	e.POST("/users", user.CreateUser)

	e.Static("/sessions", "assets")
	e.GET("/sessions", user.Index)
	e.GET("/sessions/new", user.NewSession)
	e.POST("/sessions", user.CreateSession)

	//e.GET("/facebook/oauth2", RenderFacebook)
	//e.GET("/facebook/callback", &controllersFacebook)

	e.Logger.Fatal(e.Start(":8080"))
}
