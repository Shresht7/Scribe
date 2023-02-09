package examples

import (
	"fmt"

	"github.com/Shresht7/Scribe/markdown"
	"github.com/Shresht7/Scribe/scribe"
)

func ExampleMarkdownDocument() {
	// Create a new document
	document := markdown.NewDocument()

	// Append some text nodes to the document
	document.AppendChild(
		markdown.Heading(1, "Markdown Document"),
		markdown.Paragraph(
			scribe.Text("This is a test document."),
		),
	)

	// Print the document to the console
	fmt.Println(document)

	// Output:
	// # Markdown Document
	//
	// This is a test document.
}
