package scribe

import (
	"strings"
)

//* DOCUMENT *//

// Document is just a collection of text nodes separated by a newline.
type Document struct {
	strings.Builder
	NodeContainer
}

// Instantiate a new document
func NewDocument() *Document {

	// Instantiate a new document with default options
	document := &Document{}
	document.WithSeparator("\n")

	// Return the document instance
	return document

}

// String returns the string representation of the document.
func (document *Document) String() string {

	// Reset the string builder
	document.Builder.Reset()

	// Return an empty string if there are no nodes
	if len(document.Nodes) == 0 {
		return ""
	}

	// Write the nodes to the string builder
	last := len(document.Nodes) - 1
	for _, node := range document.Nodes[:last] {
		// Write the node and the separator
		document.Builder.WriteString(node.String() + document.Separator)
	}
	// Write the last node without a separator
	document.Builder.WriteString(document.Nodes[last].String())

	// Return the string representation of the document
	return document.Builder.String()

}
