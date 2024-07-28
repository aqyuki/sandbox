package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.GET("/", Root)
	app.Start(":8080")
}

func Root(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World\n")
}
