package scribe

import (
	"strings"
)

//* NODE TEXT *//

// NodeText is a node containing text.
// It is a special type of ParentNode that joins the child nodes with a separator.
// The default separator is a single whitespace.
// The separator can be changed by calling the Join() method.
type NodeText struct {
	Separator string
	ParentNode
}

// Instantiate a new NodeText with the given contents
func Text(contents ...string) *NodeText {
	// Generate a slice of NodeLiteral from the given contents
	Nodes := []Node{}
	for _, content := range contents {
		Nodes = append(Nodes, &NodeLiteral{Value: content})
	}

	// Return the NodeText instance
	return &NodeText{
		ParentNode: ParentNode{
			Nodes: Nodes,
		},
	}
}

// Joins the text nodes with the given separator.
func (text *NodeText) Join(separator string) *NodeText {
	text.Separator = separator
	return text
}

// Implement the Node interface for NodeText
func (text *NodeText) String() string {
	res := []string{}
	for _, node := range text.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, text.Separator)
}
