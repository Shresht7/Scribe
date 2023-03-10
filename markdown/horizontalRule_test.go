package markdown

import (
	"fmt"
	"testing"
)

func TestHorizontalRule(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		hr          string
		expected    string
	}{
		{
			description: "Empty horizontal rule",
			hr:          HorizontalRule('-', 0).String(),
			expected:    "---",
		},
		{
			description: "Horizontal rule with 3 dashes",
			hr:          HorizontalRule('-', 3).String(),
			expected:    "---",
		},
		{
			description: "Horizontal rule with 5 dashes",
			hr:          HorizontalRule('-', 5).String(),
			expected:    "-----",
		},
		{
			description: "Horizontal rule with 3 stars",
			hr:          HorizontalRule('*', 3).String(),
			expected:    "***",
		},
		{
			description: "Horizontal rule with 5 stars",
			hr:          HorizontalRule('*', 5).String(),
			expected:    "*****",
		},
		{
			description: "Horizontal rule with 3 underscores",
			hr:          HorizontalRule('_', 3).String(),
			expected:    "___",
		},
		{
			description: "Horizontal rule with 5 underscores",
			hr:          HorizontalRule('_', 5).String(),
			expected:    "_____",
		},
		{
			description: "--- is the default horizontal rule",
			hr:          HorizontalRule('-', 0).String(),
			expected:    "---",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		if testCase.hr != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.hr)
		}
	}

}

func ExampleHorizontalRule() {
	md := Paragraph()

	md.AppendChild(Text("This is a paragraph with a horizontal rule below it."))
	md.AppendChild(HorizontalRule('-', 3))
	md.AppendChild(Text("This is a paragraph with a horizontal rule above it."))

	fmt.Println(md)

	// Output:
	// This is a paragraph with a horizontal rule below it.
	// ---
	// This is a paragraph with a horizontal rule above it.
}
