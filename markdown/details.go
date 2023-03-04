package markdown

import (
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

//* DETAILS *//

// A details block hides its contents until the user clicks on the
// summary element. This is useful for hiding large amounts of
// information that the user can choose to view.
type NodeDetails struct {
	Summary string
	scribe.NodeText
}

// Create a new details block
func Details(summary string, contents ...any) *NodeDetails {
	details := &NodeDetails{
		Summary: summary,
	}
	details.WithSeparator(" ")
	details.AppendChild(contents...)
	return details
}

// Implement the Node interface for details blocks
func (n *NodeDetails) String() string {
	res := []string{}

	// Open the details block
	res = append(res, "<details>")

	// Add the summary
	res = append(res, "\n<summary>"+n.Summary+"</summary>\n")

	// Add the contents
	res = append(res, n.NodeText.String())

	// Close the details block
	res = append(res, "\n</details>")

	// Return the details block as a string
	return strings.Join(res, "\n")
}

// Write the details block to the document
func (doc *MarkdownDocument) WriteDetails(summary string, contents ...any) *MarkdownDocument {
	doc.AppendChild(Details(summary, contents...))
	return doc
}
