package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed data/*
var files embed.FS

func main() {
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Printf("[file] %s\n", path)
		}
		return nil
	})
}
