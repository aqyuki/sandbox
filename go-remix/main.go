package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Static("/", "frontend/build/client")

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusNotFound && !strings.HasPrefix(c.Request().URL.Path, "/api") {
				c.Logger().Error(err)
				c.File("frontend/build/client/index.html")
			}
		} else {
			app.DefaultHTTPErrorHandler(err, c)
		}
	}

	app.Start(":8080")
}
