# cloudavenue_iam_user (Data Source)

The user data source allows you to read users in Cloud Avenue.

## Schema

### Optional

- `id` (String) The ID of the user. Ensure that one and only one attribute from this collection is set : `name`, `id`.
- `name` (String) The name of the user. Ensure that one and only one attribute from this collection is set : `name`, `id`.

### Read-Only

- `role_name` (String) The role assigned to the user.
