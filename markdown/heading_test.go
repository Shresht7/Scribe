package markdown

import (
	"fmt"
	"testing"
)

func TestHeading(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		heading     string
		expected    string
	}{
		{
			description: "Empty heading",
			heading:     Heading(1, ""),
			expected:    "# ",
		},
		{
			description: "Heading with text",
			heading:     NewHeading(1, NewText("Heading Level One")).String(),
			expected:    "# Heading Level One",
		},
		{
			description: "Second level heading with text",
			heading:     NewHeading(2, NewText("Heading Level Two")).String(),
			expected:    "## Heading Level Two",
		},
		{
			description: "Third level heading with text",
			heading:     NewHeading(3, NewText("Heading Level Three")).String(),
			expected:    "### Heading Level Three",
		},
		{
			description: "Fourth level heading with text",
			heading:     NewHeading(4, NewText("Heading Level Four")).String(),
			expected:    "#### Heading Level Four",
		},
		{
			description: "Fifth level heading with text",
			heading:     NewHeading(5, NewText("Heading Level Five")).String(),
			expected:    "##### Heading Level Five",
		},
		{
			description: "Sixth level heading with text",
			heading:     NewHeading(6, NewText("Heading Level Six")).String(),
			expected:    "###### Heading Level Six",
		},
		{
			description: "Should not render a heading with a level greater than 6",
			heading:     NewHeading(7, NewText("No such thing as a Heading Level Seven")).String(),
			expected:    "###### No such thing as a Heading Level Seven",
		},
		{
			description: "Should not render a heading with a level less than 1",
			heading:     NewHeading(0, NewText("No such thing as a Heading Level Zero")).String(),
			expected:    "# No such thing as a Heading Level Zero",
		},
		{
			description: "Heading with multiple text nodes",
			heading: NewHeading(1,
				NewText("This is sentence one."),
				NewText("This is sentence two."),
			).String(),
			expected: "# This is sentence one. This is sentence two.",
		},
		{
			description: "Heading with multiple text nodes and a separator",
			heading: func() string {
				h := NewHeading(1,
					NewText("This is sentence one."),
					NewText("This is sentence two."),
				)
				h.WithSeparator("\t")
				return h.String()
			}(),
			expected: "# This is sentence one.\tThis is sentence two.",
		},
		{
			description: "Heading with text and inline elements",
			heading: NewHeading(1,
				NewBold(NewText("This is some bold text")),
				NewText("This is a test heading."),
				NewText("This is some random text"),
			).String(),
			expected: "# **This is some bold text** This is a test heading. This is some random text",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		// Compare the actual and expected results
		if testCase.heading != testCase.expected {
			t.Errorf("TestHeading: %s: Expected %q, got %q", testCase.description, testCase.expected, testCase.heading)
		}
	}

}

func ExampleHeading() {
	// Create a new heading
	heading := Heading(1, "This is a test heading.")

	// Print the heading
	fmt.Println(heading)

	// Output:
	// # This is a test heading.
}

func ExampleH2() {
	// Create a new heading
	heading := H2("This is a test heading.")

	// Print the heading
	fmt.Println(heading)

	// Output:
	// ## This is a test heading.
}
