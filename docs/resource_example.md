# cloudavenue_iam_user (Resource)

The user resource allows you to manage local users in Cloud Avenue.

## Schema

### Required

- `name` (String) (ForceNew) The name of the user.
- `password` (String, Sensitive) The user's password. This value is never returned on read.
- `role_name` (String) The role assigned to the user.

### Optional

- `vdc_group` (String, Deprecated) (ForceNew) vDC group name. This can be an existing vDC group or a new one. This allows you to isolate your vDC.
VMs of vDCs which belong to the same vDC group can communicate together.

!!! warning "**Attribute deprecated**"

     Remove the `vdc_group` attribute configuration as it replaced by the resource [`cloudavenue_vdc_group`](https://registry.terraform.io/providers/orange-cloudavenue/cloudavenue/latest/docs/resources/vdc_group) and the attribute will be removed in the version [`v0.12.0`](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/milestone/4) of the provider. See the [GitHub issue](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/448) for more information.

### Read-Only

- `id` (String) The ID of the user.
