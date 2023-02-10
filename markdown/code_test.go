package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
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
			code:        Code(scribe.Text("")).String(),
			expected:    "``",
		},
		{
			description: "Code with text",
			code:        Code(scribe.Text("Code Text")).String(),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes",
			code:        Code(scribe.Text("Code"), scribe.Text("Text")).String(),
			expected:    "`Code Text`",
		},
		{
			description: "Code with multiple text nodes and a separator",
			code: func() string {
				c := Code(scribe.Text("Code"), scribe.Text("Text"))
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
