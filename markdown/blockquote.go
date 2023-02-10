package markdown

import "github.com/Shresht7/Scribe/scribe"

//* BLOCKQUOTE *//

// NodeBlockQuote represents a blockquote.
type NodeBlockQuote struct {
	scribe.NodeText
}

// Instantiate a new blockquote with the given contents
func BlockQuote(contents ...scribe.Node) *NodeBlockQuote {
	// Create a new blockquote
	blockquote := &NodeBlockQuote{}
	blockquote.WithSeparator(" ")

	// Append the contents to the blockquote
	blockquote.AppendChild(contents...)

	// Return the blockquote
	return blockquote
}

// Implement the Node interface for NodeBlockQuote
func (blockquote *NodeBlockQuote) String() string {
	return "> " + blockquote.NodeText.String()
}
