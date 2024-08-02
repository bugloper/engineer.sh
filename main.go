package main

import (
	"engineer.sh/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	rootHandler := handlers.RootHandler{}
	e.GET("/", rootHandler.HandlerRootPage)

	e.Logger.Fatal(e.Start(":4000"))
}
