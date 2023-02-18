package markdown

import (
	"fmt"
	"testing"
)

func TestDetails(t *testing.T) {

	// Test Cases
	cases := []struct {
		description string
		details     string
		expected    string
	}{
		{
			description: "Simple Details",
			details:     Details("description", Text("This is a simple details block")).String(),
			expected:    "<details>\n\n<summary>description</summary>\n\nThis is a simple details block\n\n</details>",
		},
	}

	// Run the test cases
	for _, c := range cases {
		if c.details != c.expected {
			t.Errorf("TestDetails(%s): expected %s, got %s", c.description, c.expected, c.details)
		}
	}

}

func ExampleDetails() {
	// Create a new details block
	details := Details("description", Text("This is a simple details block"))

	// Print the details block
	fmt.Println(details)

	// Output:
	// <details>
	//
	// <summary>description</summary>
	//
	// This is a simple details block
	//
	// </details>
}
