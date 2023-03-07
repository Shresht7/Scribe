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
			paragraph:   NewParagraph(NewText("")).String(),
			expected:    "",
		},
		{
			description: "Paragraph with text",
			paragraph:   NewParagraph(NewText("This is a test paragraph.")).String(),
			expected:    "This is a test paragraph.",
		},
		{
			description: "Paragraph with multiple text nodes",
			paragraph: NewParagraph(
				NewText("This is the first sentence."),
				NewText("This is the second sentence."),
			).String(),
			expected: "This is the first sentence.\nThis is the second sentence.",
		},
		{
			description: "Paragraph with multiple text nodes and a separator",
			paragraph: NewParagraph(
				NewText("This is the first sentence."),
				NewText("This is the second sentence."),
			).WithSeparator("\t").String(),
			expected: "This is the first sentence.\tThis is the second sentence.",
		},
		{
			description: "Paragraph with text and inline elements",
			paragraph: NewParagraph(
				NewBold(NewText("This is some bold text")),
				NewText("This is some plain text"),
				NewItalic(NewText("This is some italic text")),
			).String(),
			expected: "**This is some bold text**\nThis is some plain text\n*This is some italic text*",
		},
		{
			description: "Paragraph with text and block elements",
			paragraph: NewParagraph(
				NewText("This is some plain text"),
				NewBlockQuote(NewText("This is a blockquote")),
				NewText("This is some plain text"),
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
	paragraph := NewParagraph(
		NewText("This is a test paragraph."),
	)

	// Print the paragraph
	fmt.Println(paragraph)

	// Output: This is a test paragraph.
}

func ExampleParagraph_advanced() {
	// Create a new paragraph
	paragraph := NewParagraph(
		NewText("This is the first sentence."),
		NewBold(NewText("This is some bold text")),
		NewText("This is the second sentence."),
	)

	// Print the paragraph
	fmt.Println(paragraph)

	// Output:
	// This is the first sentence.
	// **This is some bold text**
	// This is the second sentence.
}
