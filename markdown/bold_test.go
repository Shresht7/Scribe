package markdown

import (
	"fmt"
	"testing"
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
			bold:        Bold(""),
			expected:    "****",
		},
		{
			description: "Bold with text",
			bold:        Bold("Bold Text"),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes",
			bold:        NewBold(NewText("Bold"), NewText("Text")).String(),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes and a separator",
			bold: func() string {
				b := NewBold(NewText("Bold"), NewText("Text"))
				b.WithSeparator("-")
				return b.String()
			}(),
			expected: "**Bold-Text**",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.bold != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.bold)
		}
	}

}

func ExampleBold() {
	// Create a new bold text
	bold := Bold("Bold Text")

	// Print the bold text
	fmt.Println(bold)

	// Output:
	// **Bold Text**
}

func ExampleBold_code() {
	// Create a new bold text
	bold := NewBold(
		NewCode(
			NewText("Bold Text"),
		),
	)

	// Print the bold text
	fmt.Println(bold)

	// Output:
	// **`Bold Text`**
}
