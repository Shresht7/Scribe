package scribe

import (
	"fmt"
	"testing"
)

//* NODE TEXT *//

func TestText(t *testing.T) {

	// Test Cases
	testCases := []struct {
		name      string
		input     []any
		separator string
		expected  string
	}{
		{
			name:     "Empty",
			input:    []any{},
			expected: "",
		},
		{
			name:     "Single",
			input:    []any{"Hello, world!"},
			expected: "Hello, world!",
		},
		{
			name:     "Multiple",
			input:    []any{"Hello,", "world!"},
			expected: "Hello, world!",
		},
		{
			name:      "Multiple with separator",
			input:     []any{"One", "Two", "Three"},
			separator: "---",
			expected:  "One---Two---Three",
		},
		{
			name:     "Multiple with empty strings",
			input:    []any{"One", "", "Three"},
			expected: "One  Three",
		},
		{
			name:      "Multiple with newline strings",
			input:     []any{"One", "Two", "Three"},
			separator: "\n",
			expected:  "One\nTwo\nThree",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			// Create a new text node
			text := NewText(testCase.input...)

			// Join the text if a separator is provided
			if testCase.separator != "" {
				text.WithSeparator(testCase.separator)
			}

			// Check the output
			if text.String() != testCase.expected {
				t.Errorf("%s: want %s, got %s", testCase.name, testCase.expected, text.String())
			}
		})
	}

}

func ExampleText() {

	// Create a new text node
	text := NewText("Hello,", "world!")

	// Print the string representation of the node
	fmt.Println(text)

	// Output: Hello, world!

}
