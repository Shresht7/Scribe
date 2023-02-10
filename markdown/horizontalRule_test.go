package markdown

import "testing"

func TestHorizontalRule(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		hr          string
		expected    string
	}{
		{
			description: "Empty horizontal rule",
			hr:          HorizontalRule('-', 0).String(),
			expected:    "---",
		},
		{
			description: "Horizontal rule with 3 dashes",
			hr:          HorizontalRule('-', 3).String(),
			expected:    "---",
		},
		{
			description: "Horizontal rule with 5 dashes",
			hr:          HorizontalRule('-', 5).String(),
			expected:    "-----",
		},
		{
			description: "Horizontal rule with 3 stars",
			hr:          HorizontalRule('*', 3).String(),
			expected:    "***",
		},
		{
			description: "Horizontal rule with 5 stars",
			hr:          HorizontalRule('*', 5).String(),
			expected:    "*****",
		},
		{
			description: "Horizontal rule with 3 underscores",
			hr:          HorizontalRule('_', 3).String(),
			expected:    "___",
		},
		{
			description: "Horizontal rule with 5 underscores",
			hr:          HorizontalRule('_', 5).String(),
			expected:    "_____",
		},
		{
			description: "--- is the default horizontal rule",
			hr:          HorizontalRule('-', 0).String(),
			expected:    "---",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.hr != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.hr)
		}
	}

}
