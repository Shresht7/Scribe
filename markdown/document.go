package markdown

import (
	"bytes"

	"github.com/Shresht7/Scribe/scribe"
)

//* DOCUMENT *//

// MarkdownDocument is a document that can be rendered as Markdown
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

// Implement the Node interface for MarkdownDocument
func (document *MarkdownDocument) String() string {
	// Reset the buffer to empty it because we don't want
	// the previous contents to be included in the new string
	document.buf.Reset()

	// Iterate over the nodes
	for _, node := range document.Nodes {
		// Add the node to the buffer
		document.buf.WriteString(node.String() + document.Separator)
	}

	// Return the string
	return document.buf.String()
}
