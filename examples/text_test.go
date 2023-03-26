package examples

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Shresht7/Scribe/scribe"
)

func ExampleDocument() {
	// Create a new Document
	document := scribe.NewDocument()

	// Write some text to the document
	document.WriteText("Hello World")

	// Write some more text to the document
	document.WriteText("This", "is", "a", "test", "document")

	// Print the document
	fmt.Println(document)

	// Output:
	// Hello World
	// This is a test document
}

// ExampleTextDocument demonstrates how to create a text document.
func ExampleDocument_append() {
	// Create a new document
	document := scribe.NewDocument()

	// Append some text nodes to the document
	document.AppendChild(
		scribe.NewText("Hello, world!"),
		scribe.NewText("This", "is", "a", "test", "document.").WithSeparator("---"),
	)

	// Print the document to the console
	fmt.Println(document)

	// Output:
	// Hello, world!
	// This---is---a---test---document.
}

// ExampleTextDocument demonstrates how to create a text document using a string builder.
func ExampleDocument_stringBuilder() {
	// Create a new document
	document := scribe.NewDocument()

	builder := new(strings.Builder)
	builder.WriteString("This is generated ")
	builder.WriteString("by a string builder!!!")

	// Append some text nodes to the document
	document.AppendChild(
		scribe.NewText("Example: String Builder"),
		scribe.NewText("\n---\n"),
		builder, // <== The string builder implements the Node interface!
		scribe.NewText("\n---\n"),
	)

	// Print the document to the console
	fmt.Println(document)

	// Output:
	// Example: String Builder
	//
	// ---
	//
	// This is generated by a string builder!!!
	//
	// ---
}

// ExampleTextDocument demonstrates how to create a text document using a string builder and templates.
func ExampleDocument_templates() {
	// Create a new document
	document := scribe.NewDocument()

	// Append some text nodes to the document
	document.AppendChild(
		scribe.NewText("Example: String Builder + Templates"),
	)

	// Create a template and execute it
	tmp := template.Must(template.New("test").Parse("This is {{.}}!!!"))
	// The document implements the io.Writer interface
	tmp.Execute(document, "generated by a string builder in a template")

	// Print the document to the console
	fmt.Println(document)

	// Output:
	// Example: String Builder + Templates
	// This is generated by a string builder in a template!!!
}

// ExampleTextDocument demonstrates how to write a text document to a io.Writer.
func ExampleDocument_writeToStdout() {
	// Create a new document
	document := scribe.NewDocument()

	// Append some text nodes to the document
	document.AppendChild(
		scribe.NewText("Example: Using io.Writer (os.Stdout)"),
	)

	// write the document to the writer
	fmt.Fprintln(os.Stdout, document)

	// Output:
	// Example: Using io.Writer (os.Stdout)
}
