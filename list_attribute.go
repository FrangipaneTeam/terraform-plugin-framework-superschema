// code generated by go generate - look at type_attribute.go.tmpl for source file
package superschema

import (
	"context"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ Attribute = ListAttribute{}

type ListAttribute struct {
	Common     *schemaR.ListAttribute
	Resource   *schemaR.ListAttribute
	DataSource *schemaD.ListAttribute
}

// IsResource returns true if the attribute is a resource attribute.
func (s ListAttribute) IsResource() bool {
	return s.Resource != nil || s.Common != nil
}

// IsDataSource returns true if the attribute is a data source attribute.
func (s ListAttribute) IsDataSource() bool {
	return s.DataSource != nil || s.Common != nil
}

//nolint:dupl
func (s ListAttribute) GetResource(ctx context.Context) schemaR.Attribute {
	var (
		common   schemaR.ListAttribute
		resource schemaR.ListAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.Resource != nil {
		resource = *s.Resource
	}

	a := schemaR.ListAttribute{
		Required:            computeIsRequired(common, resource),
		Optional:            computeIsOptional(common, resource),
		Computed:            computeIsComputed(common, resource),
		Sensitive:           computeIsSensitive(common, resource),
		MarkdownDescription: computeMarkdownDescription(common, resource),
		Description:         computeDescription(common, resource),
		DeprecationMessage:  computeDeprecationMessage(common, resource),
		ElementType:         common.ElementType,
	}

	if s.Resource != nil {
		if s.Resource.ElementType != nil {
			a.ElementType = s.Resource.ElementType
		}
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, resource.Validators...)
	a.PlanModifiers = append(a.PlanModifiers, common.PlanModifiers...)
	a.PlanModifiers = append(a.PlanModifiers, resource.PlanModifiers...)

	defaultVDescription := ""
	if s.Resource != nil {
		if s.Resource.Default != nil {
			a.Default = s.Resource.Default
			defaultVDescription = s.Resource.Default.MarkdownDescription(ctx)
		}
	}

	a.MarkdownDescription = genResourceAttrDescription(ctx, a.MarkdownDescription, defaultVDescription, a.Validators, a.PlanModifiers)
	return a
}

//nolint:dupl
func (s ListAttribute) GetDataSource(ctx context.Context) schemaD.Attribute {
	var (
		common     schemaR.ListAttribute
		dataSource schemaD.ListAttribute
	)

	if s.Common != nil {
		common = *s.Common
	}

	if s.DataSource != nil {
		dataSource = *s.DataSource
	}

	a := schemaD.ListAttribute{
		Required:            computeIsRequired(common, dataSource),
		Optional:            computeIsOptional(common, dataSource),
		Computed:            computeIsComputed(common, dataSource),
		Sensitive:           computeIsSensitive(common, dataSource),
		MarkdownDescription: computeMarkdownDescription(common, dataSource),
		Description:         computeDescription(common, dataSource),
		DeprecationMessage:  computeDeprecationMessage(common, dataSource),
		ElementType:         common.ElementType,
	}

	if s.DataSource != nil {
		if s.DataSource.ElementType != nil {
			a.ElementType = s.DataSource.ElementType
		}
	}

	a.Validators = append(a.Validators, common.Validators...)
	a.Validators = append(a.Validators, dataSource.Validators...)

	a.MarkdownDescription = genDataSourceAttrDescription(ctx, a.MarkdownDescription, a.Validators)

	return a
}
