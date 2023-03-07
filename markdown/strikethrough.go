package markdown

import "github.com/Shresht7/Scribe/scribe"

//* STRIKETHROUGH *//

// StrikeThrough is a strikethrough text.
type NodeStrikeThrough struct {
	scribe.NodeText
}

// Instantiate a new strikethrough text with the given contents
func NewStrikeThrough(contents ...any) *NodeStrikeThrough {
	// Create a new strikethrough text
	strikethrough := &NodeStrikeThrough{}
	strikethrough.WithSeparator(" ")

	// Append the contents to the strikethrough text
	strikethrough.AppendChild(contents...)

	// Return the strikethrough text
	return strikethrough
}

// Implement the Node interface for NodeStrikeThrough
func (bold *NodeStrikeThrough) String() string {
	return "~~" + bold.NodeText.String() + "~~"
}

func StrikeThrough(contents ...any) string {
	return NewStrikeThrough(contents...).String()
}
