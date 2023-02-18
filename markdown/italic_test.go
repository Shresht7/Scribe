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
			italic:      Italic(Text("")).String(),
			expected:    "**",
		},
		{
			description: "Italic with text",
			italic:      Italic(Text("Italic Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes",
			italic:      Italic(Text("Italic"), Text("Text")).String(),
			expected:    "*Italic Text*",
		},
		{
			description: "Italic with multiple text nodes and a separator",
			italic: func() string {
				i := Italic(Text("Italic"), Text("Text"))
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
	italic := Italic(Text("Italic Text"))
	// Print the italic
	fmt.Println(italic)
	// Output: *Italic Text*
}
