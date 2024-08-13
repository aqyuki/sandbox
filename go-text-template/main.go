package main

import (
	"flag"
	"os"
	"text/template"
)

const templateString = `package {{.PackageName}}`

var (
	packageName = flag.String("package", "main", "package name")
)

type Inventory struct {
	PackageName string
}

func main() {
	tmpl, err := template.New("sample").Parse(templateString)
	if err != nil {
		panic(err)
	}

	inventory := Inventory{
		PackageName: *packageName,
	}

	if err := tmpl.Execute(os.Stdout, inventory); err != nil {
		panic(err)
	}
}

func init() {
	flag.Parse()
}
