package superschema

import "fmt"

const (
	warningDeprecated = "\n\n ~> **Attribute deprecated** "
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

	message := warningDeprecated

	switch {
	case d.Renamed:
		if d.TargetAttributeName == "" {
			return ""
		}
		message += fmt.Sprintf("Rename the `%s` attribute to `%s` ", d.FromAttributeName, d.TargetAttributeName)
	case d.Removed:
		message += fmt.Sprintf("Remove the `%s` attribute configuration ", d.FromAttributeName)

		if d.TargetResourceName != "" {
			if d.LinkToResourceDoc != "" {
				message += fmt.Sprintf("as it replaced by the resource [`%s`](%s) ", d.TargetResourceName, d.LinkToResourceDoc)
			} else {
				message += fmt.Sprintf("as it replaced by the resource `%s` ", d.TargetResourceName)
			}
		}
	default:
		return ""
	}

	if d.LinkToMilestone != "" {
		message += fmt.Sprintf(", it will be be removed in the version [`%s`](%s) of the provider", d.TargetRelease, d.LinkToMilestone)
	} else {
		message += fmt.Sprintf(", it will be be removed in the version `%s` of the provider", d.TargetRelease)
	}

	if d.LinkToIssue != "" {
		message += fmt.Sprintf(". See the [GitHub issue](%s) for more information.", d.LinkToIssue)
	}

	return message
}

// GetDeprecationMessage returns the deprecation message for the attribute.
func (d *Deprecated) GetDeprecationMessage() string {
	return d.DeprecationMessage
}

// GetMarkdownDeprecationMessage returns the markdown deprecation message for the attribute.
func (d *Deprecated) GetMarkdownDeprecationMessage() string {
	return d.computeDeprecatedDocumentation()
}
