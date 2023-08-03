# Getting Started
The following is a technical documentation on how to use superschema in your Terraform project.

## Installation

For installing the superschema, you can use the `go get` command:

```bash
go get github.com/FrangipaneTeam/terraform-plugin-framework-superschema@latest
```

## How to use it

Add to you golang imports :

```go
import (
  superschema "github.com/FrangipaneTeam/terraform-plugin-framework-superschema"
)
```
Create, for example, a function witch define the superschema of your resource :

```go
/*
userSchema

This function is used to create the schema for the user resource and datasource.
*/
func userSchema() superschema.Schema {
	return superschema.Schema{
		Common: superschema.SchemaDetails{
			MarkdownDescription: "The user",
		},
		Resource: superschema.SchemaDetails{
			MarkdownDescription: "resource allows you to manage local users in Cloud Avenue.",
		},
		DataSource: superschema.SchemaDetails{
			MarkdownDescription: "data source allows you to read users in Cloud Avenue.",
		},
		Attributes: map[string]superschema.Attribute{
			"id": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The ID of the user.",
				},
				Resource: &schemaR.StringAttribute{
					Computed: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.UseStateForUnknown(),
					},
				},
				DataSource: &schemaD.StringAttribute{
					Optional: true,
					Computed: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(path.MatchRoot("name"), path.MatchRoot("id")),
					},
				},
			},
			"name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The name of the user.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplace(),
					},
				},
				DataSource: &schemaD.StringAttribute{
					Optional: true,
					Computed: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(path.MatchRoot("name"), path.MatchRoot("id")),
					},
				},
			},
			"role_name": superschema.StringAttribute{
				Common: &schemaR.StringAttribute{
					MarkdownDescription: "The role assigned to the user.",
				},
				Resource: &schemaR.StringAttribute{
					Required: true,
				},
				DataSource: &schemaD.StringAttribute{
					Computed: true,
				},
			},
			"password": superschema.StringAttribute{
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "The user's password. This value is never returned on read.",
					Required:            true,
					Sensitive:           true,
				},
			},
			"vdc_group": superschema.StringAttribute{
				Deprecated: &superschema.Deprecated{
					DeprecationMessage:                "Remove the vdc_group attribute configuration as it replaced by the resource cloudavenue_vdc_group and the attribute will be removed in the version 0.12.0 of the provider.",
					ComputeMarkdownDeprecationMessage: true,
					Removed:                           true,
					FromAttributeName:                 "vdc_group",
					TargetRelease:                     "v0.12.0",
					LinkToMilestone:                   "https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/milestone/4",
					TargetResourceName:                "cloudavenue_vdc_group",
					LinkToResourceDoc:                 "https://registry.terraform.io/providers/orange-cloudavenue/cloudavenue/latest/docs/resources/vdc_group",
					LinkToIssue:                       "https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/448",
				},
				Resource: &schemaR.StringAttribute{
					MarkdownDescription: "This can be an existing vDC group or a new one. This allows you to isolate your vDC.\n" +
						"VMs of vDCs which belong to the same vDC group can communicate together.",
					Optional: true,
					PlanModifiers: []planmodifier.String{
						stringplanmodifier.RequiresReplaceIfConfigured(),
					},
				},
			},
		},
	}
}
```

Look at the doc example for the generated documentation of this example :

* [Resource documentation](resource_example.md)
* [DataSource documentation](datasource_example.md)

