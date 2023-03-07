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
			blockquote:  BlockQuote(""),
			expected:    "> ",
		},
		{
			description: "Blockquote with text",
			blockquote:  BlockQuote("Blockquote Text"),
			expected:    "> Blockquote Text",
		},
		{
			description: "Blockquote with multiple text nodes",
			blockquote:  BlockQuote(NewText("Blockquote"), NewText("Text")),
			expected:    "> Blockquote\n> Text",
		},
		{
			description: "Blockquote with multiple text nodes and a separator",
			blockquote: func() string {
				b := NewBlockQuote(NewText("Blockquote"), NewText("Text"))
				b.WithSeparator("\t")
				return b.String()
			}(),
			expected: "> Blockquote\t> Text",
		},
		{
			description: "Blockquote with multiple text nodes and empty lines",
			blockquote: func() string {
				b := NewBlockQuote("Blockquote", "", NewBold("Text"))
				return b.String()
			}(),
			expected: "> Blockquote\n> \n> **Text**",
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
	blockquote := BlockQuote("Blockquote Text")

	// Print the blockquote
	fmt.Println(blockquote)

	// Output:
	// > Blockquote Text
}

func ExampleBlockQuote_multipleTextNodes() {
	// Create a new blockquote
	blockquote := NewBlockQuote(
		NewText("Blockquote"),
		NewBold("Text"),
	)

	// Print the blockquote
	fmt.Println(blockquote)

	// Output:
	// > Blockquote
	// > **Text**
}
