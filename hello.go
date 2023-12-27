package demo

import (
	"github.com/labstack/echo/v4"
)

type Message struct {
	Message string `json:"message"`
}

func Hello(c echo.Context) error {
	// m := Message{"Hello, World!"}
	panic("oops")
	// return c.JSON(http.StatusOK, m)
}
