SuperSchema allows you to write the Terraform schema for resources and datasources in a single definition, along with a common field that enables you to define default values. SuperSchema is compatible with [tfplugindocs](github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs).

This is a try to solve these issues :

* Don't repeat yourself : common fields applied on resources and datasources.
* Auto format attributes markdown description with validators and plan modifiers descriptions, default values...
