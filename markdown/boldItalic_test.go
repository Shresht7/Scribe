package markdown

import (
	"fmt"
	"testing"
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
			boldItalic:  BoldItalic(Text("")).String(),
			expected:    "******",
		},
		{
			description: "Bold italic with text",
			boldItalic:  BoldItalic(Text("Bold Italic Text")).String(),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes",
			boldItalic:  BoldItalic(Text("Bold"), Text("Italic"), Text("Text")).String(),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes and a separator",
			boldItalic: func() string {
				b := BoldItalic(Text("Bold"), Text("Italic"), Text("Text"))
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

func ExampleBoldItalic() {
	// Create a new bold italic
	boldItalic := BoldItalic(Text("Bold Italic Text"))

	// Print the bold italic
	fmt.Println(boldItalic)

	// Output:
	// ***Bold Italic Text***
}
