package markdown

import (
	"fmt"
	"testing"
)

func TestParagraph(t *testing.T) {
	// Create test cases
	testCases := []struct {
		description string
		paragraph   string
		expected    string
	}{
		{
			description: "Empty paragraph",
			paragraph:   Paragraph(Text("")).String(),
			expected:    "",
		},
		{
			description: "Paragraph with text",
			paragraph:   Paragraph(Text("This is a test paragraph.")).String(),
			expected:    "This is a test paragraph.",
		},
		{
			description: "Paragraph with multiple text nodes",
			paragraph: Paragraph(
				Text("This is the first sentence."),
				Text("This is the second sentence."),
			).String(),
			expected: "This is the first sentence.\nThis is the second sentence.",
		},
		{
			description: "Paragraph with multiple text nodes and a separator",
			paragraph: Paragraph(
				Text("This is the first sentence."),
				Text("This is the second sentence."),
			).WithSeparator("\t").String(),
			expected: "This is the first sentence.\tThis is the second sentence.",
		},
		{
			description: "Paragraph with text and inline elements",
			paragraph: Paragraph(
				Bold(Text("This is some bold text")),
				Text("This is some plain text"),
				Italic(Text("This is some italic text")),
			).String(),
			expected: "**This is some bold text**\nThis is some plain text\n*This is some italic text*",
		},
		{
			description: "Paragraph with text and block elements",
			paragraph: Paragraph(
				Text("This is some plain text"),
				BlockQuote(Text("This is a blockquote")),
				Text("This is some plain text"),
			).String(),
			expected: "This is some plain text\n> This is a blockquote\nThis is some plain text",
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		// Compare the actual and expected results
		if testCase.paragraph != testCase.expected {
			t.Errorf("TestParagraph: %s\nExpected:\n%q\ngot:\n%q\n\n", testCase.description, testCase.expected, testCase.paragraph)
		}
	}

}

func ExampleParagraph_simple() {
	// Create a new paragraph
	paragraph := Paragraph(
		Text("This is a test paragraph."),
	)

	// Print the paragraph
	fmt.Println(paragraph)

	// Output: This is a test paragraph.
}

func ExampleParagraph_advanced() {
	// Create a new paragraph
	paragraph := Paragraph(
		Text("This is the first sentence."),
		Bold(Text("This is some bold text")),
		Text("This is the second sentence."),
	)

	// Print the paragraph
	fmt.Println(paragraph)

	// Output:
	// This is the first sentence.
	// **This is some bold text**
	// This is the second sentence.
}
