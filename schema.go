package superschema

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

//go:generate go run template.go
var (
	reOneOf = regexp.MustCompile(`\\"(\S+)\\"`)
	words   = regexp.MustCompile(`(\w+)`)
)

const (
	useStateForUnknown    = "useStateForUnknownModifier"
	requireReplace        = "requiresReplaceIfModifier"
	validatorOneOf        = "oneOfValidator"
	validatorExactlyOneOf = "ExactlyOneOfValidator"

	forceNewDesc = "(ForceNew)"
)

type Schema struct {
	Deprecated Deprecated
	Common     SchemaDetails
	Resource   SchemaDetails
	DataSource SchemaDetails
	Attributes Attributes
}

type SchemaDetails struct {
	MarkdownDescription string
	DeprecationMessage  string
}

func (s Schema) GetResource(ctx context.Context) schemaR.Schema {
	if s.Resource.MarkdownDescription != "" {
		s.Common.MarkdownDescription = addToDescription(s.Common.MarkdownDescription, s.Resource.MarkdownDescription)
	}

	if s.Resource.DeprecationMessage != "" {
		s.Common.DeprecationMessage = addToDescription(s.Common.MarkdownDescription, s.Resource.DeprecationMessage)
	}

	if s.Deprecated.DeprecationMessage != "" {
		s.Common.DeprecationMessage = addToDescription(s.Common.MarkdownDescription, s.Deprecated.DeprecationMessage)
	}

	return schemaR.Schema{
		MarkdownDescription: s.Common.MarkdownDescription,
		DeprecationMessage:  s.Common.DeprecationMessage,
		Attributes:          s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
	}
}

func (s Schema) GetDataSource(ctx context.Context) schemaD.Schema {
	if s.DataSource.MarkdownDescription != "" {
		s.Common.MarkdownDescription = addToDescription(s.Common.MarkdownDescription, s.DataSource.MarkdownDescription)
	}

	if s.DataSource.DeprecationMessage != "" {
		s.Common.DeprecationMessage = addToDescription(s.Common.MarkdownDescription, s.DataSource.DeprecationMessage)
	}

	if s.Deprecated.DeprecationMessage != "" {
		s.Common.DeprecationMessage = addToDescription(s.Common.MarkdownDescription, s.Deprecated.DeprecationMessage)
	}

	return schemaD.Schema{
		MarkdownDescription: s.Common.MarkdownDescription,
		DeprecationMessage:  s.Common.DeprecationMessage,
		Attributes:          s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
	}
}

// appendToDescription appends the given description to the existing one.
func addToDescription(description, toAdd string) string {
	if toAdd == "" {
		return description
	}
	if description == "" {
		return strings.TrimLeft(toAdd, " ")
	}
	return strings.TrimRight(description, " ") + " " + strings.TrimLeft(toAdd, " ")
}

// addOneOfToDescription reformat OneOf validator description.
func addOneOfToDescription(oneof, description string) string {
	params := reOneOf.FindAllStringSubmatch(oneof, -1)
	desc := ""
	for i, p := range params {
		desc += fmt.Sprintf("`%s`", p[1])
		if i < len(params)-1 {
			desc += ", "
		}
	}
	newD := description
	newD += "Value must be one of : " + desc + "."
	return newD
}

// addOnlyOneToDescription reformat OneOf validator description.
func addOnlyOneToDescription(onlyO, description string) string {
	p := strings.Split(onlyO, ":")
	if len(p) <= 1 {
		return description
	}
	params := words.FindAllStringSubmatch(p[1], -1)
	desc := ""
	for i, p := range params {
		desc += fmt.Sprintf("`%s`", p[1])
		if i < len(params)-1 {
			desc += ", "
		}
	}
	newD := "Ensure that one and only one attribute from this collection is set : " + desc + "."
	return newD
}

// addToDescriptionWithDot add missing dots to description.
func addToDescriptionWithDot(description, toAdd string) string {
	newD := addEndDot(description)
	if toAdd == "" {
		return description
	}
	newD += fmt.Sprintf(" %s", capitalize(toAdd))
	newD = addEndDot(newD)
	return newD
}

// addEndDot adds a dot to the end of the description if it is missing.
func addEndDot(description string) string {
	if !strings.HasSuffix(description, ".") {
		return description + "."
	}

	return description
}

// getType returns the type of the given variable as a string.
func getType(myvar interface{}) string {
	t := reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

// capitalize returns the given string with the first letter capitalized.
func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// updateValidatorsDescription update description with validator description.
func updateValidatorsDescription[D validator.Describer](ctx context.Context, validators []D) string {
	description := ""
	for _, v := range validators {
		toAdd := v.MarkdownDescription(ctx)
		if toAdd == "" {
			continue
		}
		name := getType(v)
		switch name {
		case validatorOneOf:
			description = addOneOfToDescription(toAdd, description)
		case validatorExactlyOneOf:
			description = addOnlyOneToDescription(toAdd, description)
		default:
			description = addToDescription(description, capitalize(toAdd))
		}
		description = addEndDot(description)
	}
	return description
}

func updatePlanModifierDescription[D planmodifier.Describer](ctx context.Context, description string, planmodifiers []D) string {
	for _, pm := range planmodifiers {
		name := getType(pm)

		toAdd := pm.MarkdownDescription(ctx)
		if toAdd == "" {
			continue
		}

		switch name {
		case requireReplace:
			description = addToDescription(forceNewDesc, description)

		case useStateForUnknown:
			continue

		default:
			description = addToDescriptionWithDot(description, toAdd)
		}
	}
	description = addEndDot(description)
	return description
}

func genResourceAttrDescription[V validator.Describer, P planmodifier.Describer](ctx context.Context, description, defaultVDescription, deprecatedDescription string, validators []V, planmodifiers []P) string {
	pmDescription := updatePlanModifierDescription(ctx, description, planmodifiers)
	validatorDescription := updateValidatorsDescription(ctx, validators)

	description = pmDescription
	if validatorDescription != "" {
		description = addToDescriptionWithDot(description, validatorDescription)
	}
	if defaultVDescription != "" {
		description = addToDescriptionWithDot(description, defaultVDescription)
	}
	if deprecatedDescription != "" {
		description = addToDescriptionWithDot(description, deprecatedDescription)
	}

	description = addEndDot(description)
	return description
}

func genDataSourceAttrDescription[V validator.Describer](ctx context.Context, description, deprecatedDescription string, validators []V) string {
	validatorDescription := updateValidatorsDescription(ctx, validators)
	if validatorDescription != "" {
		description = addToDescriptionWithDot(description, validatorDescription)
	}
	if deprecatedDescription != "" {
		description = addToDescriptionWithDot(description, deprecatedDescription)
	}
	description = addEndDot(description)
	return description
}
