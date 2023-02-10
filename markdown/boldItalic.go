package markdown

import "github.com/Shresht7/Scribe/scribe"

//* BOLD ITALIC *//

// NodeBoldItalic represents bold italicized text.
type NodeBoldItalic struct {
	scribe.NodeText
}

// Instantiate a new bold italic text with the given contents
func BoldItalic(contents ...scribe.Node) *NodeBoldItalic {
	// Create a new bold italic
	boldItalic := &NodeBoldItalic{}
	boldItalic.WithSeparator(" ")

	// Append the contents to the bold italic
	boldItalic.AppendChild(contents...)

	// Return the bold italic
	return boldItalic
}

// Implement the Node interface for NodeBoldItalic
func (boldItalic *NodeBoldItalic) String() string {
	return "***" + boldItalic.NodeText.String() + "***"
}
