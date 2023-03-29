package superschema

import (
	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type superAttribute interface {
	schemaR.Attribute
	schemaD.Attribute
}

func computeIsComputed[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsComputed() || resourceOrDatasource.IsComputed()
}

func computeIsOptional[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsOptional() || resourceOrDatasource.IsOptional()
}

func computeIsRequired[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsRequired() || resourceOrDatasource.IsRequired()
}

func computeIsSensitive[X, T superAttribute](common X, resourceOrDatasource T) bool {
	return common.IsSensitive() || resourceOrDatasource.IsSensitive()
}

func computeMarkdownDescription[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetMarkdownDescription(), resourceOrDatasource.GetMarkdownDescription())
}

func computeDescription[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetDescription(), resourceOrDatasource.GetDescription())
}

func computeDeprecationMessage[X, T superAttribute](common X, resourceOrDatasource T) string {
	return addToDescription(common.GetDeprecationMessage(), resourceOrDatasource.GetDeprecationMessage())
}
