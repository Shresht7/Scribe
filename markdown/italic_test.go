package markdown

import (
	"fmt"
	"testing"
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
			italic:      NewItalic(NewText("")).String(),
			expected:    "**",
		},
		{
			description: "Italic with text",
			italic:      NewItalic(NewText("Italic Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes",
			italic:      NewItalic(NewText("Italic"), NewText("Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes and a separator",
			italic: func() string {
				i := NewItalic(NewText("Italic"), NewText("Text"))
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

func ExampleItalic() {
	// Create an italic
	italic := NewItalic(NewText("Italic Text"))
	// Print the italic
	fmt.Println(italic)
	// Output: *Italic Text*
}
