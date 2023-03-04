package markdown

import (
	"fmt"
	"testing"
)

func TestFencedBlock(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string // Description of the test case
		fence       string // Fenced block string
		expected    string // Expected fenced block string
	}{
		{
			description: "Empty",
			fence:       FencedBlock("").String(),
			expected:    "---\n\n---",
		},
		{
			description: "With content",
			fence:       FencedBlock("Content").String(),
			expected:    "---\nContent\n---",
		},
		{
			description: "With multiple lines of content",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").String(),
			expected:    "---\nLine 1\nLine 2\nLine 3\n---",
		},
		{
			description: "With metadata",
			fence:       FencedBlock("Content").WithMetadata("key1", "key2").String(),
			expected:    "---key1 key2\nContent\n---",
		},
		{
			description: "With multiple lines of content and metadata",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").WithMetadata("key1", "key2").String(),
			expected:    "---key1 key2\nLine 1\nLine 2\nLine 3\n---",
		},
		{
			description: "With custom fence",
			fence:       FencedBlock("Content").WithFence("~~~").String(),
			expected:    "~~~\nContent\n~~~",
		},
		{
			description: "With custom fence and metadata",
			fence:       FencedBlock("Content").WithFence("~~~").WithMetadata("key1", "key2").String(),
			expected:    "~~~key1 key2\nContent\n~~~",
		},
		{
			description: "With custom fence and multiple lines of content",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").WithFence("~~~").String(),
			expected:    "~~~\nLine 1\nLine 2\nLine 3\n~~~",
		},
		{
			description: "With custom fence, multiple lines of content and metadata",
			fence:       FencedBlock("Line 1\nLine 2\nLine 3").WithFence("~~~").WithMetadata("key1", "key2").String(),
			expected:    "~~~key1 key2\nLine 1\nLine 2\nLine 3\n~~~",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		if testCase.fence != testCase.expected {
			t.Errorf("Expected:\n%s\n===\ngot:\n%s", testCase.expected, testCase.fence)
		}
	}

}

func ExampleFencedBlock() {
	// Create a fenced block with content
	fencedBlock := FencedBlock("Content")

	// Output the fenced block
	fmt.Println(fencedBlock)

	// Output:
	// ---
	// Content
	// ---
}
