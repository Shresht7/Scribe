package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestHeading(t *testing.T) {

	// Create test cases
	testCases := []struct {
		description string
		heading     string
		expected    string
	}{
		{
			description: "Empty heading",
			heading:     Heading(1, scribe.Text("")).String(),
			expected:    "# ",
		},
		{
			description: "Heading with text",
			heading:     Heading(1, scribe.Text("Heading Level One")).String(),
			expected:    "# Heading Level One",
		},
		{
			description: "Second level heading with text",
			heading:     Heading(2, scribe.Text("Heading Level Two")).String(),
			expected:    "## Heading Level Two",
		},
		{
			description: "Third level heading with text",
			heading:     Heading(3, scribe.Text("Heading Level Three")).String(),
			expected:    "### Heading Level Three",
		},
		{
			description: "Fourth level heading with text",
			heading:     Heading(4, scribe.Text("Heading Level Four")).String(),
			expected:    "#### Heading Level Four",
		},
		{
			description: "Fifth level heading with text",
			heading:     Heading(5, scribe.Text("Heading Level Five")).String(),
			expected:    "##### Heading Level Five",
		},
		{
			description: "Sixth level heading with text",
			heading:     Heading(6, scribe.Text("Heading Level Six")).String(),
			expected:    "###### Heading Level Six",
		},
		{
			description: "Should not render a heading with a level greater than 6",
			heading:     Heading(7, scribe.Text("No such thing as a Heading Level Seven")).String(),
			expected:    "###### No such thing as a Heading Level Seven",
		},
		{
			description: "Should not render a heading with a level less than 1",
			heading:     Heading(0, scribe.Text("No such thing as a Heading Level Zero")).String(),
			expected:    "# No such thing as a Heading Level Zero",
		},
		{
			description: "Heading with multiple text nodes",
			heading: Heading(1,
				scribe.Text("This is sentence one."),
				scribe.Text("This is sentence two."),
			).String(),
			expected: "# This is sentence one. This is sentence two.",
		},
		{
			description: "Heading with multiple text nodes and a separator",
			heading: func() string {
				h := Heading(1,
					scribe.Text("This is sentence one."),
					scribe.Text("This is sentence two."),
				)
				h.WithSeparator("\t")
				return h.String()
			}(),
			expected: "# This is sentence one.\tThis is sentence two.",
		},
		{
			description: "Heading with text and inline elements",
			heading: Heading(1,
				Bold("This is some bold text"),
				scribe.Text("This is a test heading."),
				scribe.Text("This is some random text"),
			).String(),
			expected: "# **This is some bold text** This is a test heading. This is some random text",
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		// Compare the actual and expected results
		if testCase.heading != testCase.expected {
			t.Errorf("TestHeading: %s: Expected %q, got %q", testCase.description, testCase.expected, testCase.heading)
		}
	}

}
