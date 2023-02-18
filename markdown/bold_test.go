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
			bold:        Bold(Text("")).String(),
			expected:    "****",
		},
		{
			description: "Bold with text",
			bold:        Bold(Text("Bold Text")).String(),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes",
			bold:        Bold(Text("Bold"), Text("Text")).String(),
			expected:    "**Bold Text**",
		},
		{
			description: "Bold with multiple text nodes and a separator",
			bold: func() string {
				b := Bold(Text("Bold"), Text("Text"))
				b.WithSeparator(" ")
				return b.String()
			}(),
			expected: "**Bold Text**",
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
	bold := Bold(Text("Bold Text"))

	// Print the bold text
	fmt.Println(bold)

	// Output:
	// **Bold Text**
}

func ExampleBold_code() {
	// Create a new bold text
	bold := Bold(
		Code(
			Text("Bold Text"),
		),
	)

	// Print the bold text
	fmt.Println(bold)

	// Output:
	// **`Bold Text`**
}
