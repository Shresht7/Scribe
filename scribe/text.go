package scribe

//* NODE TEXT *//

// NodeText is a node containing text.
// It is a special type of ParentNode that joins the child nodes
// with a single whitespace separator.
type NodeText struct {
	NodeParent
}

// Instantiate a new NodeText with the given contents
func Text(contents ...any) *NodeText {
	// Create a new text node
	text := &NodeText{
		NodeParent: NodeParent{
			Separator: " ",
		},
	}

	// Append the contents to the text node
	text.AppendChild(contents...)

	// Return the text node
	return text
}

// Write a text node to the document
func (doc *Document) WriteText(contents ...any) *Document {
	doc.AppendChild(Text(contents...))
	return doc
}
