package markdown

import (
	"fmt"
	"testing"
)

func TestImage(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		image       string
		expected    string
	}{
		{
			description: "Test image with alt and url",
			image:       Image("Alt", "URL").String(),
			expected:    "![Alt](URL)",
		},
		{
			description: "Test image with real alt and url",
			image:       Image("Google", "https://www.google.com").String(),
			expected:    "![Google](https://www.google.com)",
		},
		{
			description: "Test image with empty alt and url",
			image:       Image("", "").String(),
			expected:    "![]()",
		},
		{
			description: "Test image with relative url",
			image:       Image("Relative", "/relative").String(),
			expected:    "![Relative](/relative)",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.image != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.image)
		}
	}

}

func ExampleImage() {
	// Create an image
	image := Image("Alt", "URL")
	// Print the image
	fmt.Println(image)
	// Output: ![Alt](URL)
}
