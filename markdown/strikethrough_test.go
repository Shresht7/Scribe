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
			strikethrough: NewStrikeThrough(NewText("")).String(),
			expected:      "~~~~",
		},
		{
			name:          "Single",
			strikethrough: NewStrikeThrough(NewText("Hello")).String(),
			expected:      "~~Hello~~",
		},
		{
			name:          "Multiple",
			strikethrough: NewStrikeThrough(NewText("Hello"), NewText("World")).String(),
			expected:      "~~Hello World~~",
		},
		{
			name: "Multiple with separator",
			strikethrough: func() string {
				s := NewStrikeThrough(NewText("Hello"), NewText("World"))
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
	strikethrough := NewStrikeThrough(NewText("Hello"))
	// Print the strikethrough
	fmt.Println(strikethrough)
	// Output: ~~Hello~~
}
