package scribe

//* NODE TEXT *//

// NodeText is a node containing text.
// It is a special type of ParentNode that joins the child nodes with a separator.
// The default separator is a single whitespace.
// The separator can be changed by calling the WithSeparator() method.
type NodeText struct {
	*ParentNode
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
		&ParentNode{
			Separator: " ",
			Nodes:     Nodes,
		},
	}
}
