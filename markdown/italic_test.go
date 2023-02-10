package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestItalic(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		italic      string
		expected    string
	}{
		{
			description: "Empty italic",
			italic:      Italic(scribe.Text("")).String(),
			expected:    "**",
		},
		{
			description: "Italic with text",
			italic:      Italic(scribe.Text("Italic Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes",
			italic:      Italic(scribe.Text("Italic"), scribe.Text("Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes and a separator",
			italic: func() string {
				i := Italic(scribe.Text("Italic"), scribe.Text("Text"))
				i.WithSeparator(" ")
				return i.String()
			}(),
			expected: "*Italic Text*",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.italic != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.italic)
		}
	}

}
