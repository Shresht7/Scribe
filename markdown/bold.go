package markdown

import "github.com/Shresht7/Scribe/scribe"

//* BOLD *//

// Bold is a bold text.
type NodeBold struct {
	scribe.NodeText
}

// Instantiate a new bold text with the given contents
func Bold(contents ...scribe.Node) *NodeBold {
	// Create a new bold text
	bold := &NodeBold{}
	bold.WithSeparator(" ")

	// Append the contents to the bold text
	bold.AppendChild(contents...)

	// Return the bold text
	return bold
}

// Implement the Node interface for NodeBold
func (bold *NodeBold) String() string {
	return "**" + bold.NodeText.String() + "**"
}
