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
				scribe.Text("This is a test paragraph."),
				scribe.Text("This is another test paragraph."),
			).String(),
			expected: "This is a test paragraph. This is another test paragraph.",
		},
		{
			description: "Paragraph with multiple text nodes and a separator",
			paragraph: Paragraph(
				scribe.Text("This is a test paragraph."),
				scribe.Text("This is another test paragraph."),
			).WithSeparator("\t").String(),
			expected: "This is a test paragraph.\tThis is another test paragraph.",
		},
		// {
		// 	description: "Paragraph with text and inline elements",
		// 	paragraph: Paragraph(
		// 		Bold("This is some bold text"),
		// 		scribe.Text("This is a test paragraph."),
		// 		scribe.Text("This is another test paragraph."),
		// 	).String(),
		// 	expected: "**This is some bold text** This is a test paragraph. This is another test paragraph.",
		// },
	}

	// Run the test cases
	for _, testCase := range testCases {
		// Compare the actual and expected results
		if testCase.paragraph != testCase.expected {
			t.Errorf("TestParagraph: %s: Expected %q, got %q", testCase.description, testCase.expected, testCase.paragraph)
		}
	}

}
