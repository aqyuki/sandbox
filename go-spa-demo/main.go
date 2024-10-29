package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed dist
var embeddedFiles embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      true,
		Filesystem: getFileSystem("dist"),
	}))

	e.Group("assets").Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}), middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: getFileSystem("dist/assets"),
	}))

	e.Group("api").GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}

func getFileSystem(path string) http.FileSystem {
	fs, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fs)
}
