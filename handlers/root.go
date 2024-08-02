package handlers

import (
	"engineer.sh/views/home"
	"github.com/labstack/echo/v4"
)

type RootHandler struct {
}

func (h RootHandler) HandlerRootPage(c echo.Context) error {
	return render(c, home.Root())
}
