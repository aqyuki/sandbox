package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed static/*
var embedFS embed.FS

func main() {
	sub, err := fs.Sub(embedFS, "static")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Printf("sub: %v\n", sub)

	sub, err = fs.Sub(embedFS, "static/assets")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Printf("sub: %v\n", sub)

	f, err := sub.Open("asset.txt")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Printf("content: %s\n", buf[:n])

	sub, err = fs.Sub(embedFS, "static/scripts")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Printf("sub: %v\n", sub)

	sub, err = fs.Sub(embedFS, "static/styles")
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
		return
	}
	fmt.Printf("sub: %v\n", sub)
}
