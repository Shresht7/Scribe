package markdown

import (
	"fmt"
	"testing"
)

func TestStrikeThrough(t *testing.T) {

	// Create test cases
	testCases := []struct {
		name          string
		strikethrough string
		expected      string
	}{
		{
			name:          "Empty",
			strikethrough: StrikeThrough(Text("")).String(),
			expected:      "~~~~",
		},
		{
			name:          "Single",
			strikethrough: StrikeThrough(Text("Hello")).String(),
			expected:      "~~Hello~~",
		},
		{
			name:          "Multiple",
			strikethrough: StrikeThrough(Text("Hello"), Text("World")).String(),
			expected:      "~~Hello World~~",
		},
		{
			name: "Multiple with separator",
			strikethrough: func() string {
				s := StrikeThrough(Text("Hello"), Text("World"))
				s.WithSeparator("_")
				return s.String()
			}(),
			expected: "~~Hello_World~~",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		if testCase.strikethrough != testCase.expected {
			t.Errorf("Expected %s, got %s", testCase.expected, testCase.strikethrough)
		}
	}

}

func ExampleStrikeThrough() {
	// Create a strikethrough
	strikethrough := StrikeThrough(Text("Hello"))
	// Print the strikethrough
	fmt.Println(strikethrough)
	// Output: ~~Hello~~
}
