package main

import (
	"fmt"
	"os"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

func main() {
	name := os.Args[1]

	repositoryName := fmt.Sprintf("%sRepository", name)

	// generate repository code
	generateRepositoryCode(repositoryName)
}

func generateRepositoryCode(name string) {
	// generate new repository
	f := NewFile("repository")

	// generate lower camel case repository name
	lowerCamelName := strcase.ToLowerCamel(name)

	// generate new repository struct
	f.Type().Id(lowerCamelName).Struct()

	// generate new repository interface
	f.Type().Id(name).Interface()

	// generate new repository func
	f.Func().Id(fmt.Sprintf("New%s", name)).Call().Id(name).Block(
		Return().Op("&").Id(lowerCamelName).Values(),
	)

	fp, err := os.Create(fmt.Sprintf("repository/%s_repository.go", strcase.ToSnake(name)))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	fp.WriteString(fmt.Sprintf("%#v", f))
}