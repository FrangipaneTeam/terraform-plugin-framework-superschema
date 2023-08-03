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

## DeprecationMessage

The `DeprecationMessage` is the special field of the superschema. Is used to manage deprecation of attributes.
List of possible values:

```go
    // DeprecationMessage is the message to display in the CLI when the user
 // attempts to use the deprecated attribute.
 // This field is required.
 DeprecationMessage string

 // MarkdownDeprecationMessage is the message to display in the Documentation portal
 // when the user attempts to use the deprecated attribute.
 // This field is optional if ComputeMarkdownDeprecationMessage is false.
 MarkdownDeprecationMessage string

 // ComputeMarkdownDeprecationMessage is a flag to indicate whether the MarkdownDeprecationMessage
 // should be computed from the parameters of the Deprecated struct.
 ComputeMarkdownDeprecationMessage bool

 // Renamed is a flag to indicate whether the attribute has been renamed.
 // Removed is a flag to indicate whether the attribute has been removed.
 // One of these fields must be true.
 Renamed, Removed bool

 // FromAttributeName is the name of the attribute that has been deprecated.
 // This field is required if ComputeMarkdownDeprecationMessage is true.
 FromAttributeName string

 // TargetAttributeName is the name of the attribute that replaces the deprecated attribute.
 // TargetResourceName is the name of the resource that replaces the deprecated attribute.
 // These fields are optional if the attribute has been removed.
 TargetAttributeName, TargetResourceName string

 // TargetRelease is the release version in which the attribute was deprecated. (e.g. v1.0.0).
 // This field is Required.
 TargetRelease string
 // LinkToIssue is the link to the GitHub issue that describes the deprecation.
 // This field is optional.
 LinkToIssue string
 // LinkToResourceDoc is the link to the terraform documentation for the resource that replaces the deprecated attribute.
 // This field is optional.
 LinkToResourceDoc string
 // LinkToMilestone is the link to the GitHub milestone that describes the release in which the attribute was deprecated.
 // This field is optional.
 LinkToMilestone string
```
