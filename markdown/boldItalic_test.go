package markdown

import (
	"fmt"
	"testing"
)

func TestBoldItalic(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		boldItalic  string
		expected    string
	}{
		{
			description: "Empty bold italic",
			boldItalic:  BoldItalic(""),
			expected:    "******",
		},
		{
			description: "Bold italic with text",
			boldItalic:  BoldItalic("Bold Italic Text"),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes",
			boldItalic:  NewBoldItalic(NewText("Bold"), NewText("Italic"), NewText("Text")).String(),
			expected:    "***Bold Italic Text***",
		},
		{
			description: "Bold italic with multiple text nodes and a separator",
			boldItalic: func() string {
				b := NewBoldItalic(NewText("Bold"), NewText("Italic"), NewText("Text"))
				b.WithSeparator("-")
				return b.String()
			}(),
			expected: "***Bold-Italic-Text***",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		if testCase.boldItalic != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.boldItalic)
		}
	}

}

func ExampleBoldItalic() {
	// Create a new bold italic
	boldItalic := BoldItalic("Bold Italic Text")

	// Print the bold italic
	fmt.Println(boldItalic)

	// Output:
	// ***Bold Italic Text***
}
