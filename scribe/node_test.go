package scribe

import (
	"fmt"
	"testing"
)

//* NODE LITERAL *//

func TestNodeLiteral(t *testing.T) {

	// Test the String() method

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
			literal:     &NodeLiteral{"Hello World"},
			want:        "Hello World",
		},
		{
			description: "Literal with Integer Value",
			literal:     NewLiteral(1),
			want:        "1",
		},
		{
			description: "Literal with Boolean Value",
			literal:     NewLiteral(true),
			want:        "true",
		},
		{
			description: "Literal with Slice Value",
			literal:     NewLiteral([]int{1, 2, 3}),
			want:        "[1 2 3]",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := testCase.literal.String()
			if got != testCase.want {
				t.Errorf("%s: got %q, want %q", testCase.description, got, testCase.want)
			}
		})
	}

}

func TestNodeLiteral_Write(t *testing.T) {

	// Test the Write() method

	// Test Cases
	testCases := []struct {
		description string
		literal     *NodeLiteral
		input       string
		want        string
	}{
		{
			description: "Empty Literal",
			literal:     &NodeLiteral{},
			input:       "Hello World",
			want:        "Hello World",
		},
		{
			description: "Literal with Value",
			literal:     &NodeLiteral{"Hello"},
			input:       " World",
			want:        "Hello World",
		},
		{
			description: "Literal with Integer Value",
			literal:     NewLiteral(1),
			input:       " 2 3",
			want:        "1 2 3",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			_, err := fmt.Fprint(testCase.literal, testCase.input)
			if err != nil {
				t.Errorf("%s: %s", testCase.description, err)
			}

			got := testCase.literal.String()
			if got != testCase.want {
				t.Errorf("%s: got %q, want %q", testCase.description, got, testCase.want)
			}
		})
	}

}

func ExampleNodeLiteral() {

	// Create a new NodeLiteral
	node := &NodeLiteral{"Hello"}
	fmt.Println(node)

	// or use the Literal() function
	node = NewLiteral("World")
	fmt.Println(node)

	// Output:
	// Hello
	// World

}

//* NODE CONTAINER *//

func TestNodeContainer(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		parent      *NodeContainer
		want        string
	}{
		{
			description: "Empty Parent",
			parent:      &NodeContainer{},
			want:        "",
		},
		{
			description: "Parent with Children",
			parent: &NodeContainer{
				Separator: " ",
				Nodes: []Node{
					NewLiteral("Hello"),
					&NodeLiteral{"World"},
				},
			},
			want: "Hello World",
		},
		{
			description: "Parent with other Containers as Children",
			parent: &NodeContainer{
				Separator: " ",
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
			want: "One Two-Three-Four Five|Six",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			got := testCase.parent.String()
			if got != testCase.want {
				t.Errorf("%s: got %q, want %q", testCase.description, got, testCase.want)
			}
		})
	}

}

func TestNodeContainer_Write(t *testing.T) {

	// Test Cases
	testCases := []struct {
		description string
		parent      *NodeContainer
		input       string
		want        string
	}{
		{
			description: "Empty Parent",
			parent:      &NodeContainer{},
			input:       "Hello World",
			want:        "Hello World",
		},
		{
			description: "Parent with Children",
			parent: &NodeContainer{
				Separator: " ",
				Nodes: []Node{
					NewLiteral("Hello"),
					&NodeLiteral{"World"},
				},
			},
			input: "!",
			want:  "Hello World!",
		},
		{
			description: "Parent with other Containers as Children",
			parent: &NodeContainer{
				Separator: " ",
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
			input: "\n===\n",
			want:  "One Two-Three-Four Five|Six\n===\n",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			_, err := fmt.Fprint(testCase.parent, testCase.input)
			if err != nil {
				t.Errorf("%s: %s", testCase.description, err)
			}

			got := testCase.parent.String()
			if got != testCase.want {
				t.Errorf("%s: got %q, want %q", testCase.description, got, testCase.want)
			}
		})
	}

}

func ExampleNodeContainer() {

	// Create a new NodeParent
	node := &NodeContainer{
		Separator: " ",
		Nodes: []Node{
			NewLiteral("Hello"),
			&NodeLiteral{"World"},
		},
	}

	// or use the Container() function
	node = NewContainer(
		NewLiteral("Hello"),
		&NodeLiteral{"World"},
	).WithSeparator(" ")

	// Print the string representation of the node
	fmt.Println(node)

	// Output: Hello World

}
