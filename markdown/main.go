package markdown

import (
	"bytes"

	"github.com/Shresht7/Scribe/scribe"
)

//* DOCUMENT *//

// Document is a document.
type MarkdownDocument struct {
	buf *bytes.Buffer
	scribe.Document
}

// Instantiate a new document with the given contents
func NewDocument() *MarkdownDocument {
	// Create a new document
	document := scribe.NewDocument()

	// Set the separator to a double newline for Markdown
	document.WithSeparator("\n\n")

	// Return the document
	return &MarkdownDocument{
		buf:      bytes.NewBuffer([]byte{}),
		Document: *document,
	}
}

// Implement the Node interface for NodeDocument
func (document *MarkdownDocument) String() string {
	// Reset the buffer to empty it
	document.buf.Reset()

	// Iterate over the nodes
	for _, node := range document.Nodes {
		// Add the node to the buffer
		document.buf.WriteString(node.String() + document.Separator)
	}

	// Return the string
	return document.buf.String()
}
