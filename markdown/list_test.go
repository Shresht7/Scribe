package markdown

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Shresht7/Scribe/scribe"
)

func TestOrderedList(t *testing.T) {

	list := []any{
		scribe.NewText("Item 1"),
		scribe.NewText("Item 2"),
		scribe.NewText("Item 3"),
	}

	// Create test cases
	testCases := []struct {
		name     string
		list     *NodeList
		expected string
	}{
		{
			name: "Ordered List",
			list: NewOrderedList(list),
			expected: strings.Join([]string{
				"1. Item 1",
				"2. Item 2",
				"3. Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List",
			list: NewUnorderedList("*", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with -",
			list: NewUnorderedList("-", list),
			expected: strings.Join([]string{
				"- Item 1",
				"- Item 2",
				"- Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with +",
			list: NewUnorderedList("+", list),
			expected: strings.Join([]string{
				"+ Item 1",
				"+ Item 2",
				"+ Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with invalid decorator",
			list: NewUnorderedList("x", list),
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

func ExampleNewOrderedList() {
	// Create a list of nodes
	list := []any{
		NewText("Item 1"),
		NewText("Item 2"),
		NewText("Item 3"),
	}

	// Create an ordered list
	listNode := NewOrderedList(list)

	// Print the list
	fmt.Println(listNode)

	// Output:
	// 1. Item 1
	// 2. Item 2
	// 3. Item 3
}

func TestUnorderedList(t *testing.T) {

	list := []any{
		scribe.NewText("Item 1"),
		scribe.NewText("Item 2"),
		scribe.NewText("Item 3"),
	}

	// Create test cases
	testCases := []struct {
		name     string
		list     *NodeList
		expected string
	}{
		{
			name: "Ordered List",
			list: NewOrderedList(list),
			expected: strings.Join([]string{
				"1. Item 1",
				"2. Item 2",
				"3. Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List",
			list: NewUnorderedList("*", list),
			expected: strings.Join([]string{
				"* Item 1",
				"* Item 2",
				"* Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with -",
			list: NewUnorderedList("-", list),
			expected: strings.Join([]string{
				"- Item 1",
				"- Item 2",
				"- Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with +",
			list: NewUnorderedList("+", list),
			expected: strings.Join([]string{
				"+ Item 1",
				"+ Item 2",
				"+ Item 3",
			}, "\n") + "\n",
		},
		{
			name: "Unordered List with invalid decorator",
			list: NewUnorderedList("x", list),
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
	list := []any{
		NewText("Item 1"),
		NewText("Item 2"),
		NewText("Item 3"),
	}

	// Create an unordered list
	listNode := NewUnorderedList("*", list)

	// Print the list
	fmt.Println(listNode)

	// Output:
	// * Item 1
	// * Item 2
	// * Item 3
}
