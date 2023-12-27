package main

import (
	"demo"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", demo.Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
