package scribe

import (
	"fmt"
	"testing"
)

func TestNodeLiteral(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		literal     *NodeLiteral
		want        string
	}{
		{
			description: "Empty Literal",
			literal:     &NodeLiteral{},
			want:        "",
		},
		{
			description: "Literal with Value",
			literal:     &NodeLiteral{Value: "Hello World"},
			want:        "Hello World",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := testCase.literal.String()
			if got != testCase.want {
				t.Errorf("got %q, want %q", got, testCase.want)
			}
		})
	}

}

func ExampleNodeLiteral() {

	// Create a new NodeLiteral
	node := &NodeLiteral{"Hello World"}

	// or use the Literal() function
	node = Literal("Hello World")

	// Print the string representation of the node
	fmt.Println(node)

	// Output: Hello World

}

func TestNodeParent(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		parent      *NodeParent
		want        string
	}{
		{
			description: "Empty Parent",
			parent:      &NodeParent{},
			want:        "",
		},
		{
			description: "Parent with Children",
			parent: &NodeParent{
				Separator: " ",
				Nodes: []Node{
					Literal("Hello"),
					&NodeLiteral{"World"},
				},
			},
			want: "Hello World",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := testCase.parent.String()
			if got != testCase.want {
				t.Errorf("got %q, want %q", got, testCase.want)
			}
		})
	}

}

func ExampleNodeParent() {

	// Create a new NodeParent
	node := &NodeParent{
		Separator: " ",
		Nodes: []Node{
			Literal("Hello"),
			&NodeLiteral{"World"},
		},
	}

	// or use the Container() function
	node = Container(
		Literal("Hello"),
		&NodeLiteral{"World"},
	).WithSeparator(" ")

	// Print the string representation of the node
	fmt.Println(node)

	// Output: Hello World

}
