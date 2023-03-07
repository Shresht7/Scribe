package markdown

import "github.com/Shresht7/Scribe/scribe"

//* ITALIC *//

// Italic represents italicized text in Markdown.
type NodeItalic struct {
	scribe.NodeText
}

// Instantiate a new italic text with the given contents
func NewItalic(contents ...any) *NodeItalic {
	// Create a new italic
	italic := &NodeItalic{}
	italic.WithSeparator(" ")

	// Append the contents to the italic
	italic.AppendChild(contents...)

	// Return the italic
	return italic
}

// Implement the Node interface for NodeItalic
func (italic *NodeItalic) String() string {
	return "*" + italic.NodeText.String() + "*"
}

// Create a new italic text with the given contents
func Italic(contents ...any) string {
	return NewItalic(contents...).String()
}
