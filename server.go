package main

import (
	"log"
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
	user := &model.User{}
	db.First(user)
	log.Printf("----------------------------------------")
	log.Println(user.Name)
	log.Printf("----------------------------------------")

	e.Logger.Fatal(e.Start(":8080"))
}
