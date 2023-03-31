//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"html/template"
	"os"

	"github.com/iancoleman/strcase"
)

type templateInfos struct {
	TypeName string
}

//go:embed type_attribute.go.tmpl
var templateTypeAttribute string

func main() {
	fmt.Println("generating types files...")
	tA := []string{"string", "bool", "float64", "int64", "list", "list_nested", "object", "set", "set_nested", "number", "single_nested"}

	for _, t := range tA {
		infos := templateInfos{
			TypeName: strcase.ToCamel(t),
		}

		tmpl, err := template.New("template").Parse(templateTypeAttribute)
		if err != nil {
			fmt.Printf("error from template parse : %v\n", err)
			os.Exit(1)
		}

		var tpl bytes.Buffer

		errExec := tmpl.Execute(&tpl, infos)

		if errExec != nil {
			fmt.Printf("error from template execute : %v\n", errExec)
			os.Exit(1)
		}

		// format the code
		formattedContent, errFormat := format.Source(tpl.Bytes())
		if errFormat != nil {
			fmt.Printf("error from format : %v\n", errFormat)
			os.Exit(1)
		}

		errWrite := os.WriteFile(t+"_attribute.go", formattedContent, 0o644)
		if errWrite != nil {
			fmt.Printf("write to file error : %v\n", errWrite)
		}
	}
}
