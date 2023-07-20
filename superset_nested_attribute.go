// code generated by go generate - look at supertype_attribute.go.tmpl for source file
package superschema

import (
	"context"

	"github.com/FrangipaneTeam/terraform-plugin-framework-supertypes"
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ Attribute = SuperSetNestedAttribute{}

type SuperSetNestedAttribute struct {
	Common     *schemaR.SetNestedAttribute
	Resource   *schemaR.SetNestedAttribute
	DataSource *schemaD.SetNestedAttribute
	Attributes Attributes
}

// IsResource returns true if the attribute is a resource attribute.
func (s SuperSetNestedAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s SuperSetNestedAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

// GetCustomType returns the custom type of the attribute.
func (s SuperSetNestedAttribute) getCustomType(oB basetypes.ObjectTypable) basetypes.SetTypable {
	return supertypes.SetNestedType{
		SetType: basetypes.SetType{
			ElemType: oB,
		},
	}
}

//nolint:dupl
func (s SuperSetNestedAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.SetNestedAttribute
		resource schemaR.SetNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.SetNestedAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		NestedObject: schemaR.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, resourceT).(map[string]schemaR.Attribute),
		},
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, resource.Validators...)
	a.PlanModifiers = append(a.PlanModifiers, common.PlanModifiers...)
	a.PlanModifiers = append(a.PlanModifiers, resource.PlanModifiers...)

	defaultVDescription := ""

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.Resource != nil {
		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
			defaultVDescription = s.Resource.Default.MarkdownDescription(ctx)
		}
		if s.Resource.CustomType != nil {
			a.CustomType = s.Resource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		a.CustomType = s.getCustomType(a.NestedObject.Type()).(supertypes.SetNestedType)
	}

	a.MarkdownDescription = genResourceAttrDescription(ctx, a.MarkdownDescription, defaultVDescription, a.Validators, a.PlanModifiers)
	return a
}

//nolint:dupl
func (s SuperSetNestedAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.SetNestedAttribute
		dataSource schemaD.SetNestedAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.SetNestedAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		NestedObject: schemaD.NestedAttributeObject{
			Attributes: s.Attributes.process(ctx, dataSourceT).(map[string]schemaD.Attribute),
		},
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, dataSource.Validators...)

	if s.Common != nil {
		if s.Common.CustomType != nil {
			a.CustomType = s.Common.CustomType
		}
	}

	if s.DataSource != nil {
		if s.DataSource.CustomType != nil {
			a.CustomType = s.DataSource.CustomType
		}
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		a.CustomType = s.getCustomType(a.NestedObject.Type()).(supertypes.SetNestedType)
	}
	// * If user has not provided a custom type, we will use the default supertypes
	if a.CustomType == nil {
		a.CustomType = s.getCustomType(a.NestedObject.Type()).(supertypes.SetNestedType)
	}

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, a.Validators)
	return a
}
