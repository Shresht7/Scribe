package markdown

import (
	"github.com/Shresht7/Scribe/scribe"
)

//* DOCUMENT *//

// MarkdownDocument is a document that can be rendered as Markdown
type MarkdownDocument struct {
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
		Document: *document,
	}
}
