package markdown

import "github.com/Shresht7/Scribe/scribe"

//* STRIKETHROUGH *//

// NodeStrikethrough is a strikethrough text.
type NodeStrikethrough struct {
	scribe.NodeText
}

// Instantiate a new strikethrough text with the given contents
func NewStrikethrough(contents ...any) *NodeStrikethrough {
	// Create a new strikethrough text
	strikethrough := &NodeStrikethrough{}
	strikethrough.WithSeparator(" ")

	// Append the contents to the strikethrough text
	strikethrough.AppendChild(contents...)

	// Return the strikethrough text
	return strikethrough
}

// Implement the Node interface for NodeStrikeThrough
func (bold *NodeStrikethrough) String() string {
	return "~~" + bold.NodeText.String() + "~~"
}

func Strikethrough(contents ...any) string {
	return NewStrikethrough(contents...).String()
}
