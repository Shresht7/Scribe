package markdown

import (
	"fmt"
	"testing"
)

func TestBlockQuote(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		blockquote  string
		expected    string
	}{
		{
			description: "Empty blockquote",
			blockquote:  BlockQuote(Text("")).String(),
			expected:    "> ",
		},
		{
			description: "Blockquote with text",
			blockquote:  BlockQuote(Text("Blockquote Text")).String(),
			expected:    "> Blockquote Text",
		},
		{
			description: "Blockquote with multiple text nodes",
			blockquote:  BlockQuote(Text("Blockquote"), Text("Text")).String(),
			expected:    "> Blockquote\n> Text",
		},
		{
			description: "Blockquote with multiple text nodes and a separator",
			blockquote: func() string {
				b := BlockQuote(Text("Blockquote"), Text("Text"))
				b.WithSeparator("\t")
				return b.String()
			}(),
			expected: "> Blockquote\t> Text",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		if testCase.blockquote != testCase.expected {
			t.Errorf("Expected:\n%s\ngot:\n%s", testCase.expected, testCase.blockquote)
		}
	}

}

func ExampleBlockQuote() {
	// Create a new blockquote
	blockquote := BlockQuote(Text("Blockquote Text"))

	// Print the blockquote
	fmt.Println(blockquote)

	// Output:
	// > Blockquote Text
}

func ExampleBlockQuote_multipleTextNodes() {
	// Create a new blockquote
	blockquote := BlockQuote(
		Text("Blockquote"),
		Bold(Text("Text")),
	)

	// Print the blockquote
	fmt.Println(blockquote)

	// Output:
	// > Blockquote
	// > **Text**
}
