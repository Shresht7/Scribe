package markdown

import (
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

//* BLOCKQUOTE *//

// NodeBlockQuote represents a blockquote in Markdown.
type NodeBlockQuote struct {
	scribe.NodeText
}

// Instantiate a new blockquote with the given contents
func BlockQuote(contents ...scribe.Node) *NodeBlockQuote {
	// Create a new blockquote
	blockquote := &NodeBlockQuote{}
	blockquote.WithSeparator("\n")

	// Append the contents to the blockquote
	blockquote.AppendChild(contents...)

	// Return the blockquote
	return blockquote
}

// Implement the Node interface for NodeBlockQuote
func (blockquote *NodeBlockQuote) String() string {
	res := []string{}
	for _, node := range blockquote.Nodes {
		res = append(res, "> "+node.String())
	}
	return strings.Join(res, blockquote.Separator)
}
