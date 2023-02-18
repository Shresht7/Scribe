package markdown

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestOrderedList(t *testing.T) {

	list := []scribe.Node{
		scribe.Text("Item 1"),
		scribe.Text("Item 2"),
		scribe.Text("Item 3"),
	}

	// Create test cases
	testCases := []struct {
		name     string
		list     *NodeList
		expected string
	}{
		{
			name: "Ordered List",
			list: OrderedList(list),
			expected: strings.Join([]string{
				"1. Item 1",
				"2. Item 2",
				"3. Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List",
			list: UnorderedList("*", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with -",
			list: UnorderedList("-", list),
			expected: strings.Join([]string{
				"- Item 1",
				"- Item 2",
				"- Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with +",
			list: UnorderedList("+", list),
			expected: strings.Join([]string{
				"+ Item 1",
				"+ Item 2",
				"+ Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with invalid decorator",
			list: UnorderedList("x", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
	}

	// Iterate over the test cases
	for _, testCase := range testCases {
		// Run the test case
		t.Run(testCase.name, func(t *testing.T) {
			// Check if the string representation of the list is equal to the expected string
			if testCase.list.String() != testCase.expected {
				t.Errorf("want:\n%s\n\ngot:\n%s", testCase.expected, testCase.list.String())
			}
		})
	}

}

func ExampleOrderedList() {
	// Create a list of nodes
	list := []scribe.Node{
		Text("Item 1"),
		Text("Item 2"),
		Text("Item 3"),
	}

	// Create an ordered list
	listNode := OrderedList(list)

	// Print the list
	fmt.Println(listNode)

	// Output:
	// 1. Item 1
	// 2. Item 2
	// 3. Item 3
}

func TestUnorderedList(t *testing.T) {

	list := []scribe.Node{
		scribe.Text("Item 1"),
		scribe.Text("Item 2"),
		scribe.Text("Item 3"),
	}

	// Create test cases
	testCases := []struct {
		name     string
		list     *NodeList
		expected string
	}{
		{
			name: "Ordered List",
			list: OrderedList(list),
			expected: strings.Join([]string{
				"1. Item 1",
				"2. Item 2",
				"3. Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List",
			list: UnorderedList("*", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with -",
			list: UnorderedList("-", list),
			expected: strings.Join([]string{
				"- Item 1",
				"- Item 2",
				"- Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with +",
			list: UnorderedList("+", list),
			expected: strings.Join([]string{
				"+ Item 1",
				"+ Item 2",
				"+ Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with invalid decorator",
			list: UnorderedList("x", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
	}

	// Iterate over the test cases
	for _, testCase := range testCases {
		// Run the test case
		t.Run(testCase.name, func(t *testing.T) {
			// Check if the string representation of the list is equal to the expected string
			if testCase.list.String() != testCase.expected {
				t.Errorf("want:\n%s\n\ngot:\n%s", testCase.expected, testCase.list.String())
			}
		})
	}

}

func ExampleUnorderedList() {
	// Create a list of nodes
	list := []scribe.Node{
		Text("Item 1"),
		Text("Item 2"),
		Text("Item 3"),
	}

	// Create an unordered list
	listNode := UnorderedList("*", list)

	// Print the list
	fmt.Println(listNode)

	// Output:
	// * Item 1
	// * Item 2
	// * Item 3
}
