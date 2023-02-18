package markdown

import (
	"fmt"
	"testing"

	"github.com/Shresht7/Scribe/helpers"
)

func TestRenderHeader(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Widths []int
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Widths: []int{5, 3},
			Result: "| Name  | Age |",
		},
		{
			Header: []string{"Name", "Age"},
			Widths: []int{7, 3},
			Result: "| Name    | Age |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Widths: []int{7, 3, 6, 6},
			Result: "| Name    | Age | Height | Weight |",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := Table(testCase.Header, [][]string{})
		header := table.renderHeader(testCase.Widths)
		if header != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, header)
		}
	}
}

func TestRenderDivider(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Widths []int
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Widths: []int{5, 3},
			Result: "| ----- | --- |",
		},
		{
			Header: []string{"Name", "Age"},
			Widths: []int{7, 3},
			Result: "| ------- | --- |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Widths: []int{7, 3, 6, 6},
			Result: "| ------- | --- | ------ | ------ |",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := Table(testCase.Header, [][]string{})
		divider := table.renderDivider(testCase.Widths)
		if divider != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, divider)
		}
	}
}

func TestRenderRow(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Widths []int
		Row    []string
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Widths: []int{5, 3},
			Row:    []string{"Alice", "20"},
			Result: "| Alice | 20  |",
		},
		{
			Header: []string{"Name", "Age"},
			Widths: []int{7, 3},
			Row:    []string{"Alice", "20"},
			Result: "| Alice   | 20  |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Widths: []int{7, 3, 6, 6},
			Row:    []string{"Alice", "20", "5'6", "120"},
			Result: "| Alice   | 20  | 5'6    | 120    |",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := Table(testCase.Header, [][]string{})
		row := table.renderRow(testCase.Row, testCase.Widths)
		if row != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, row)
		}
	}
}

func TestTable(t *testing.T) {

	// Create a table
	table := Table([]string{"Name", "Age"}, [][]string{
		{"Alice", "20"},
		{"Bob", "30"},
	})

	// Check the table string
	if table.String() != helpers.HereDoc(`
		| Name  | Age |
		| ----- | --- |
		| Alice | 20  |
		| Bob   | 30  |
	`) {
		t.Error("Table string is not correct")
	}

}

func ExampleTable() {
	// Create a table
	table := Table([]string{"Name", "Age"}, [][]string{
		{"Alice", "20"},
		{"Bob", "30"},
	})

	// Print the table
	fmt.Println(table)

	// Output:
	// | Name  | Age |
	// | ----- | --- |
	// | Alice | 20  |
	// | Bob   | 30  |
}

// -------
// HELPERS
// -------

// helper function to check if two slices are equal
func equal[T comparable](a, b []T) bool {
	// check if lengths are equal
	if len(a) != len(b) {
		return false
	}
	// check if elements are equal
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
