package markdown

import (
	"fmt"
	"testing"
)

func TestCodeBlock(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string // Description of the test case
		codeblock   string // The code block as a string
		expected    string // The expected output
	}{
		{
			description: "Empty",
			codeblock:   CodeBlock("").String(),
			expected:    "```\n\n```",
		},
		{
			description: "With Code",
			codeblock:   CodeBlock("fmt.Printf(\"Number: %d\", 12)").String(),
			expected:    "```\nfmt.Printf(\"Number: %d\", 12)\n```",
		},
		{
			description: "With Code and Language",
			codeblock:   CodeBlock("fmt.Printf(\"Number: %d\", 12)", "go").String(),
			expected:    "```go\nfmt.Printf(\"Number: %d\", 12)\n```",
		},
		{
			description: "With Code and More Metadata",
			codeblock:   CodeBlock("fmt.Printf(\"Number: %d\", 12)").WithMetadata("go", "runnable").String(),
			expected:    "```go runnable\nfmt.Printf(\"Number: %d\", 12)\n```",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.codeblock != testCase.expected {
			t.Errorf("Expected:\n%s\n\ngot:\n%s", testCase.expected, testCase.codeblock)
		}
	}

}

func ExampleCodeBlock() {
	// Create a new code block
	codeblock := CodeBlock("fmt.Printf(\"Number: %d\", 12)", "go")

	// Print the code block
	fmt.Println(codeblock)

	// Output:
	// ```go
	// fmt.Printf("Number: %d", 12)
	// ```
}
