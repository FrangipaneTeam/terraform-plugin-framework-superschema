package superschema

import _ "github.com/iancoleman/strcase"

type schemaType int

const (
	resourceT schemaType = iota
	dataSourceT
)
