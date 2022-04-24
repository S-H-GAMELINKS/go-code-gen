package main

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {
	name := os.Args[1]

	// generate model code
	generateModelCode(name)
}

func generateModelCode(name string) {
	// generate new model
	f := NewFile("model")

	// generate new model struct
	f.Type().Id(name).Struct()

	fp, err := os.Create(fmt.Sprintf("model/%s.go", strcase.ToSnake(name)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	fp.WriteString(fmt.Sprintf("%#v", f))
}