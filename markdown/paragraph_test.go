package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
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
			paragraph:   Paragraph(scribe.Text("")).String(),
			expected:    "",
		},
		{
			description: "Paragraph with text",
			paragraph:   Paragraph(scribe.Text("This is a test paragraph.")).String(),
			expected:    "This is a test paragraph.",
		},
		{
			description: "Paragraph with multiple text nodes",
			paragraph: Paragraph(
				scribe.Text("This is the first sentence."),
				scribe.Text("This is the second sentence."),
			).String(),
			expected: "This is the first sentence.\nThis is the second sentence.",
		},
		{
			description: "Paragraph with multiple text nodes and a separator",
			paragraph: Paragraph(
				scribe.Text("This is the first sentence."),
				scribe.Text("This is the second sentence."),
			).WithSeparator("\t").String(),
			expected: "This is the first sentence.\tThis is the second sentence.",
		},
		{
			description: "Paragraph with text and inline elements",
			paragraph: Paragraph(
				Bold(scribe.Text("This is some bold text")),
				scribe.Text("This is some plain text"),
				Italic(scribe.Text("This is some italic text")),
			).String(),
			expected: "**This is some bold text**\nThis is some plain text\n*This is some italic text*",
		},
		{
			description: "Paragraph with text and block elements",
			paragraph: Paragraph(
				scribe.Text("This is some plain text"),
				BlockQuote(scribe.Text("This is a blockquote")),
				scribe.Text("This is some plain text"),
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
