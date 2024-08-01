package main

import (
	"engineer.sh/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userHandler := handlers.UserHandler{}
	e.GET("/user", userHandler.HandlerUserListing)

	e.Logger.Fatal(e.Start(":4000"))
}
