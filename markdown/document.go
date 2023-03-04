package markdown

import (
	"bytes"

	"github.com/Shresht7/Scribe/scribe"
)

//* DOCUMENT *//

// MarkdownDocument is a document that can be rendered as Markdown
type MarkdownDocument struct {
	*bytes.Buffer
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
		Buffer:   bytes.NewBuffer([]byte{}),
		Document: *document,
	}
}

// Synchronizes the document bytes buffer with the document content nodes
func (doc *MarkdownDocument) Sync() {
	// Reset the buffer to empty it because we don't want
	// the previous contents to be included in the new string
	doc.Buffer.Reset()

	// Iterate over the nodes
	for _, node := range doc.Nodes {
		// Add the node to the buffer
		str := node.String() + doc.Separator
		doc.Buffer.Write([]byte(str))
	}
}
