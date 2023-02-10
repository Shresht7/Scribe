package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestBoldItalic(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		boldItalic  string
		expected    string
	}{
		{
			description: "Empty bold italic",
			boldItalic:  BoldItalic(scribe.Text("")).String(),
			expected:    "******",
		},
		{
			description: "Bold italic with text",
			boldItalic:  BoldItalic(scribe.Text("Bold Italic Text")).String(),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes",
			boldItalic:  BoldItalic(scribe.Text("Bold"), scribe.Text("Italic"), scribe.Text("Text")).String(),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes and a separator",
			boldItalic: func() string {
				b := BoldItalic(scribe.Text("Bold"), scribe.Text("Italic"), scribe.Text("Text"))
				b.WithSeparator(" ")
				return b.String()
			}(),
			expected: "***Bold Italic Text***",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.boldItalic != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.boldItalic)
		}
	}

}
