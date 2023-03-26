package scribe

import (
	"testing"
)

//* DOCUMENT *//

func TestDocument(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		document    *Document
		want        string
	}{
		{
			description: "Empty Document",
			document:    &Document{},
			want:        "",
		},
		{
			description: "Document with Nodes",
			document: &Document{
				NodeContainer: NodeContainer{
					Separator: "\n",
					Nodes: []Node{
						NewLiteral("Hello"),
						&NodeLiteral{"World"},
					},
				},
			},
			want: "Hello\nWorld",
		},
		{
			description: "Document with nested Nodes",
			document: &Document{
				NodeContainer: NodeContainer{
					Separator: "\n",
					Nodes: []Node{
						&NodeLiteral{"One"},
						&NodeContainer{
							Separator: "-",
							Nodes: []Node{
								NewLiteral("Two"),
								NewLiteral("Three"),
								NewLiteral("Four"),
							},
						},
						NewContainer(
							NewLiteral("Five"),
							NewLiteral("Six"),
						).WithSeparator("|"),
					},
				},
			},
			want: "One\nTwo-Three-Four\nFive|Six",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := testCase.document.String()
			if got != testCase.want {
				t.Errorf("%s: got %q, want %q", testCase.description, got, testCase.want)
			}
		})
	}

}
