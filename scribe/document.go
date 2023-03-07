package scribe

//* DOCUMENT *//

// Document is just a collection of text nodes separated by a newline.
type Document struct {
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
