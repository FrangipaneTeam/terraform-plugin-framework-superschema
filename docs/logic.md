# SuperSchema logic
## superschema.Schema
The superschema.Schema type has three fields: `Common`, `Resource`, and `DataSource`. By using these fields, you can set the description header of your Terraform resource and datasource. The superschema merges the common markdown description with the resource markdown description for the Terraform resource and the common markdown description with the datasource markdown description for the Terraform datasource.

There is no need to add extra space in the markdown description. The superschema automatically adds necessary space on the merged markdown description. This same logic applies for DeprecationMessage.

## Attributes
On the attributes side, the same logic applies for the `MarkdownDescription`, `Description`, and `DeprecationMessage`. If any of the `Computed`, `Required`, `Optional`, and `Sensitive` fields is `true` in `Common`, `Resource`, or `DataSource`, the value is `true`.

**Examples:**

* If `Required` is `true` in `Common` and `false` in `Resource`, the result is a `Required: true` for the `Resource`.
* If `Required` is `false` in `Common` and `true` in `DataSource`, the result is a `Required: true` for the `DataSource`.

## Validators
The `Validators` are the sum of the `Common` and `Resource` or `DataSource` Validators. The validator markdown description is added to the markdown description of your attribute.

## PlanModifiers
The `PlanModifiers` are the sum of the `Common` and `Resource` or `DataSource` PlanModifiers. The plan modifier markdown description is added to the markdown description of your attribute with the following exceptions:

* `UseStateForUnknown`: nothing is added.
* `RequiresReplace`: `(ForceNew)` is added at the beginning of the markdown description.

## Default Value
If you have a default value, the markdown description is ended with the markdown description of this default.
