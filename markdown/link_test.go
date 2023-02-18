package markdown

import (
	"fmt"
	"testing"
)

func TestLink(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		link        string
		expected    string
	}{
		{
			description: "Test link with title and url",
			link:        Link("Title", "URL").String(),
			expected:    "[Title](URL)",
		},
		{
			description: "Test link with real title and url",
			link:        Link("Google", "https://www.google.com").String(),
			expected:    "[Google](https://www.google.com)",
		},
		{
			description: "Test link with empty title and url",
			link:        Link("", "").String(),
			expected:    "[]()",
		},
		{
			description: "Test link with relative url",
			link:        Link("Relative", "/relative").String(),
			expected:    "[Relative](/relative)",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.link != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.link)
		}
	}

}

func ExampleLink() {
	// Create a link
	link := Link("Title", "URL")
	// Print the link
	fmt.Println(link)
	// Output: [Title](URL)
}
