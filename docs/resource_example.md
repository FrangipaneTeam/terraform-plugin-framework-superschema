# cloudavenue_iam_user (Resource)

The user resource allows you to manage local users in Cloud Avenue.

## Schema

### Required

- `name` (String) (ForceNew) The name of the user.
- `password` (String, Sensitive) The user's password. This value is never returned on read.
- `role_name` (String) The role assigned to the user.

### Read-Only

- `id` (String) The ID of the user.
