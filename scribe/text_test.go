package scribe

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {

	// Create test cases
	testCases := []struct {
		name      string
		input     []string
		separator string
		expected  string
	}{
		{
			name:     "Empty",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Single",
			input:    []string{"Hello, world!"},
			expected: "Hello, world!",
		},
		{
			name:     "Multiple",
			input:    []string{"Hello,", "world!"},
			expected: "Hello, world!",
		},
		{
			name:      "Multiple with separator",
			input:     []string{"One", "Two", "Three"},
			separator: "---",
			expected:  "One---Two---Three",
		},
		{
			name:     "Multiple with empty strings",
			input:    []string{"One", "", "Three"},
			expected: "One  Three",
		},
		{
			name:      "Multiple with newline strings",
			input:     []string{"One", "Two", "Three"},
			separator: "\n",
			expected:  "One\nTwo\nThree",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Create a new text node
			text := Text(testCase.input...)

			// Join the text if a separator is provided
			if testCase.separator != "" {
				text.WithSeparator(testCase.separator)
			}

			// Check the output
			if text.String() != testCase.expected {
				t.Errorf("Expected %s, got %s", testCase.expected, text.String())
			}
		})
	}

}

func ExampleText() {
	// Create a new text node
	text := Text("Hello,", "world!")

	// Print the string representation of the node
	fmt.Println(text)

	// Output: Hello, world!
}
