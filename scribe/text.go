package scribe

//* NODE TEXT *//

// NodeText is a node containing text.
// It is a special type of NodeContainer that joins the child nodes
// with a single whitespace separator.
type NodeText struct {
	*NodeContainer
}

// Instantiate a new NodeText with the given contents
func NewText(contents ...any) *NodeText {
	// Create a new text node
	text := &NodeText{
		NodeContainer: &NodeContainer{
			Separator: " ",
		},
	}

	// Append the contents to the text node
	text.AppendChild(contents...)

	// Return the text node
	return text
}

// Write out the text node as a string
func Text(contents ...any) string {
	return NewText(contents...).String()
}

// Write a text node to the document
func (doc *Document) WriteText(contents ...any) *Document {
	doc.AppendChild(NewText(contents...))
	return doc
}

// EmptyLine text node
func EmptyLine() string {
	return NewText("").String()
}

// Write an empty line to the document
func (doc *Document) WriteEmptyLine() *Document {
	doc.AppendChild(NewText(""))
	return doc
}
