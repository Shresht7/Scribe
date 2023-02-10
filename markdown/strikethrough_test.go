package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
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
			strikethrough: StrikeThrough(scribe.Text("")).String(),
			expected:      "~~~~",
		},
		{
			name:          "Single",
			strikethrough: StrikeThrough(scribe.Text("Hello")).String(),
			expected:      "~~Hello~~",
		},
		{
			name:          "Multiple",
			strikethrough: StrikeThrough(scribe.Text("Hello"), scribe.Text("World")).String(),
			expected:      "~~Hello World~~",
		},
		{
			name: "Multiple with separator",
			strikethrough: func() string {
				s := StrikeThrough(scribe.Text("Hello"), scribe.Text("World"))
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
