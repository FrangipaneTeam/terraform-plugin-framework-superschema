package superschema

import "fmt"

const (
	warningAttributeDeprecated = "\n\n ~> **Attribute deprecated** "
	warningResourceDeprecated  = "\n\n !> **Resource deprecated** "
)

// Deprecated is a struct to describe a deprecated attribute.
type Deprecated struct {
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

	// OnlyResource and OnlyDataSource are flags to indicate whether the deprecation message should be displayed
	// only for the resource or only for the data source.
	// If not set, the deprecation message will be displayed for both.
	OnlyResource, OnlyDataSource *bool
}

func (d *Deprecated) computeDeprecatedDocumentation() string {
	if (!d.ComputeMarkdownDeprecationMessage && d.MarkdownDeprecationMessage == "") || (d.FromAttributeName == "" && d.ComputeMarkdownDeprecationMessage) || d.TargetRelease == "" {
		return ""
	}

	if d.MarkdownDeprecationMessage != "" {
		return d.MarkdownDeprecationMessage
	}

	message := warningAttributeDeprecated

	switch {
	case d.Renamed:
		if d.TargetAttributeName == "" {
			return ""
		}
		message += fmt.Sprintf("Rename the `%s` attribute to `%s`", d.FromAttributeName, d.TargetAttributeName)
	case d.Removed:
		message += fmt.Sprintf("Remove the `%s` attribute configuration", d.FromAttributeName)

		if d.TargetResourceName != "" {
			if d.LinkToResourceDoc != "" {
				message += fmt.Sprintf("as it replaced by the resource [`%s`](%s)", d.TargetResourceName, d.LinkToResourceDoc)
			} else {
				message += fmt.Sprintf("as it replaced by the resource `%s`", d.TargetResourceName)
			}
		}
	default:
		return ""
	}

	if d.LinkToMilestone != "" {
		message += fmt.Sprintf(", it will be removed in the version [`%s`](%s) of the provider", d.TargetRelease, d.LinkToMilestone)
	} else {
		message += fmt.Sprintf(", it will be removed in the version `%s` of the provider", d.TargetRelease)
	}

	if d.LinkToIssue != "" {
		message += fmt.Sprintf(". See the [GitHub issue](%s) for more information.", d.LinkToIssue)
	}

	return addEndDot(message)
}

// GetDeprecationMessage returns the deprecation message for the attribute.
func (d *Deprecated) GetDeprecationMessage() string {
	return d.DeprecationMessage
}

// GetMarkdownDeprecationMessage returns the markdown deprecation message for the attribute.
func (d *Deprecated) GetMarkdownDeprecationMessage() string {
	return d.computeDeprecatedDocumentation()
}

// DeprecatedResource is a struct to describe a deprecated resource or data source.
type DeprecatedResource struct {
	// DeprecationMessage is the message to display in the CLI when the user
	// attempts to use the deprecated resource.
	// This field is required.
	DeprecationMessage string

	// MarkdownDeprecationMessage is the message to display in the Documentation portal
	// when the user attempts to use the deprecated attribute.
	// This field is required if ComputeMarkdownDeprecationMessage is false.
	MarkdownDeprecationMessage string

	// ComputeMarkdownDeprecationMessage is a flag to indicate whether the MarkdownDeprecationMessage
	// should be computed from the parameters of the Deprecated struct.
	ComputeMarkdownDeprecationMessage bool

	// Renamed is a flag to indicate whether the resource or datasource has been renamed.
	// Removed is a flag to indicate whether the resource or datasource has been removed.
	// One of these fields must be true.
	Renamed, Removed bool

	// TargetResourceName is the name of the resource that replaces the deprecated resource or data source.
	// These fields are required if the resource or data source has been renamed and computeMarkdownDeprecationMessage is true.
	TargetResourceName string

	// TargetRelease is the release version in which the resource or datasource was deprecated. (e.g. v1.0.0).
	// This field is Required.
	TargetRelease string
	// LinkToIssue is the link to the GitHub issue that describes the deprecation.
	// This field is optional.
	LinkToIssue string
	// LinkToMigrationGuide is the link to the actual resource documentation.
	// This field is optional.
	LinkToMigrationGuide string
	// LinkToNewResourceDoc is the link to the terraform documentation for the resource that replaces the deprecated attribute.
	// This field is optional.
	LinkToNewResourceDoc string
	// LinkToMilestone is the link to the GitHub milestone that describes the release in which the attribute was deprecated.
	// This field is optional.
	LinkToMilestone string
}

func (d *DeprecatedResource) computeDeprecatedDocumentation(isResource bool) string {

	if (!d.ComputeMarkdownDeprecationMessage && d.MarkdownDeprecationMessage == "") || (d.Renamed && d.TargetResourceName == "" && d.ComputeMarkdownDeprecationMessage) || d.TargetRelease == "" {
		return ""
	}

	if d.Removed && d.Renamed {
		return ""
	}

	ressOrData := func() string {
		if isResource {
			return "resource"
		}
		return "data source"
	}()

	message := warningResourceDeprecated + d.MarkdownDeprecationMessage

	if d.ComputeMarkdownDeprecationMessage {
		switch {
		case d.Renamed:
			if d.TargetResourceName == "" {
				return ""
			}
			if d.LinkToNewResourceDoc != "" {
				message += fmt.Sprintf("The %s has renamed to [`%s`](%s)", ressOrData, d.TargetResourceName, d.LinkToNewResourceDoc)
			} else {
				message += fmt.Sprintf("The %s has renamed to `%s`", ressOrData, d.TargetResourceName)
			}
		case d.Removed:
			message += fmt.Sprintf("The %s has been removed", ressOrData)
		default:
			return ""
		}
	}

	if d.LinkToMilestone != "" {
		message += fmt.Sprintf(", it will be removed in the version [`%s`](%s) of the provider", d.TargetRelease, d.LinkToMilestone)
	} else {
		message += fmt.Sprintf(", it will be removed in the version `%s` of the provider", d.TargetRelease)
	}

	if d.LinkToIssue != "" {
		message += fmt.Sprintf(". See the [GitHub issue](%s) for more information.", d.LinkToIssue)
	}

	return addEndDot(message)
}

// GetDeprecationMessage returns the deprecation message for the attribute.
func (d *DeprecatedResource) GetDeprecationMessage() string {

	if d.LinkToMigrationGuide != "" {
		return fmt.Sprintf("%s. See the migration guide(%s) for more information.", d.DeprecationMessage, d.LinkToMigrationGuide)
	}

	return d.DeprecationMessage
}

// GetMarkdownDeprecationMessage returns the markdown deprecation message for the attribute.
func (d *DeprecatedResource) GetMarkdownDeprecationMessage(isResource bool) string {
	return d.computeDeprecatedDocumentation(isResource)
}
