package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestBlockQuote(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		blockquote  string
		expected    string
	}{
		{
			description: "Empty blockquote",
			blockquote:  BlockQuote(scribe.Text("")).String(),
			expected:    "> ",
		},
		{
			description: "Blockquote with text",
			blockquote:  BlockQuote(scribe.Text("Blockquote Text")).String(),
			expected:    "> Blockquote Text",
		},
		{
			description: "Blockquote with multiple text nodes",
			blockquote:  BlockQuote(scribe.Text("Blockquote"), scribe.Text("Text")).String(),
			expected:    "> Blockquote Text",
		},
		{
			description: "Blockquote with multiple text nodes and a separator",
			blockquote: func() string {
				b := BlockQuote(scribe.Text("Blockquote"), scribe.Text("Text"))
				b.WithSeparator(" ")
				return b.String()
			}(),
			expected: "> Blockquote Text",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.blockquote != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.blockquote)
		}
	}

}
