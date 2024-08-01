package handlers

import (
	"engineer.sh/views/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandlerUserListing(c echo.Context) error {
	return render(c, user.List())
}
