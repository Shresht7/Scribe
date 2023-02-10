package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestBold(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		bold        string
		expected    string
	}{
		{
			description: "Empty bold",
			bold:        Bold(scribe.Text("")).String(),
			expected:    "****",
		},
		{
			description: "Bold with text",
			bold:        Bold(scribe.Text("Bold Text")).String(),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes",
			bold:        Bold(scribe.Text("Bold"), scribe.Text("Text")).String(),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes and a separator",
			bold: func() string {
				b := Bold(scribe.Text("Bold"), scribe.Text("Text"))
				b.WithSeparator(" ")
				return b.String()
			}(),
			expected: "**Bold Text**",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.bold != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.bold)
		}
	}

}
