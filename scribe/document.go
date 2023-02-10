package scribe

//* DOCUMENT *//

// Document is just a collection of text nodes separated by a newline.
type Document struct {
	NodeText
}

// Instantiate a new document
func NewDocument(opts ...*Document) *Document {
	// Instantiate a new document with default options
	document := &Document{}
	document.WithSeparator("\n")

	// If options are passed, override the defaults
	if len(opts) > 0 {
		options := opts[0]

		// Override default options with user options
		if options.Separator != "" {
			document.Separator = options.Separator
		}
	}

	// Return the document instance
	return document
}
