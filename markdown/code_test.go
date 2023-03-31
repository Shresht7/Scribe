package markdown

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		code        string
		expected    string
	}{
		{
			description: "Empty code",
			code:        Code(""),
			expected:    "``",
		},
		{
			description: "Code with text",
			code:        Code("Code Text"),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes",
			code:        NewCode(NewText("Code"), NewText("Text")).String(),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes and a separator",
			code: func() string {
				c := NewCode(NewText("Code"), NewText("Text"))
				c.WithSeparator("-")
				return c.String()
			}(),
			expected: "`Code-Text`",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		if testCase.code != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.code)
		}
	}

}

func ExampleCode() {
	// Create a new code text
	code := Code("Code Text")

	// Print the code text
	fmt.Println(code)

	// Output:
	// `Code Text`
}
