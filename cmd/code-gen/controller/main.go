package main

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {
	name := os.Args[1]

	controllerName := fmt.Sprintf("%sController", name)

	// generate controller code
	generatesControllerCode(controllerName)
}

func generatesControllerCode(name string) {
	// generate new controller
	f := NewFile("controller")

	// generate lower camel case controller name
	lowerCamelName := strcase.ToLowerCamel(name)

	// generate new controller struct
	f.Type().Id(lowerCamelName).Struct()

	// generate new controller interface
	f.Type().Id(name).Interface()

	// generate new repository func
	f.Func().Id(fmt.Sprintf("New%s", name)).Call().Id(name).Block(
		Return().Op("&").Id(lowerCamelName).Values(),
	)

	fp, err := os.Create(fmt.Sprintf("controller/%s_controller.go", strcase.ToSnake(name)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	fp.WriteString(fmt.Sprintf("%#v", f))
}