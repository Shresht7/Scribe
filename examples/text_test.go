package examples

import (
	"fmt"

	"github.com/Shresht7/Scribe/scribe"
)

// ExampleTextDocument demonstrates how to create a text document.
func ExampleDocument() {
	// Create a new document
	document := scribe.NewDocument()

	// Append some text nodes to the document
	document.AppendChild(
		scribe.Text("Hello, world!"),
		scribe.Text("This", "is", "a", "test", "document.").Join("---"),
	)

	// Print the document to the console
	fmt.Println(document)

	// Output:
	// Hello, world!
	// This---is---a---test---document.
}
