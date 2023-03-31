//go:build ignore

// go generate
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"unicode"

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
		}

		var tpl bytes.Buffer

		errExec := tmpl.Execute(&tpl, infos)

		if errExec != nil {
			fmt.Printf("error from template execute : %v\n", errExec)
		}

		errWrite := os.WriteFile(t+"_attribute.go", tpl.Bytes(), 0o644)
		if errWrite != nil {
			fmt.Printf("write to file error : %v\n", errWrite)
		}
	}
}

// capitalize returns the given string with the first letter capitalized.
func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
