package markdown

import "testing"

func TestCodeBlock(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		codeblock   string
		expected    string
	}{
		{
			description: "Empty code block",
			codeblock:   CodeBlock("").String(),
			expected:    "```\n\n```",
		},
		{
			description: "Code block with text",
			codeblock:   CodeBlock("Code Block Text").String(),
			expected:    "```\nCode Block Text\n```",
		},
		{
			description: "Code block with a language",
			codeblock:   CodeBlock("Code Block Text", "go").String(),
			expected:    "```go\nCode Block Text\n```",
		},
		{
			description: "Code block with multiple languages",
			codeblock:   CodeBlock("Code Block Text", "go", "python").String(),
			expected:    "```go python\nCode Block Text\n```",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.codeblock != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.codeblock)
		}
	}

}
