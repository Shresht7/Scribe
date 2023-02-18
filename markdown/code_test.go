package markdown

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		code        string
		expected    string
	}{
		{
			description: "Empty code",
			code:        Code(Text("")).String(),
			expected:    "``",
		},
		{
			description: "Code with text",
			code:        Code(Text("Code Text")).String(),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes",
			code:        Code(Text("Code"), Text("Text")).String(),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes and a separator",
			code: func() string {
				c := Code(Text("Code"), Text("Text"))
				c.WithSeparator(" ")
				return c.String()
			}(),
			expected: "`Code Text`",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.code != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.code)
		}
	}

}

func ExampleCode() {
	// Create a new code text
	code := Code(Text("Code Text"))

	// Print the code text
	fmt.Println(code)

	// Output:
	// `Code Text`
}
