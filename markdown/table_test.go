package markdown

import (
	"fmt"
	"testing"

	helpers "github.com/Shresht7/goutils/strings"
)

func TestRenderHeader(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Rows   [][]string
		Align  map[int]string
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Rows: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
			},
			Result: "| Name  | Age |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Rows: [][]string{
				{"Alice", "20", "5'6", "120"},
				{"Bob", "30", "5'10", "150"},
			},
			Result: "| Name  | Age | Height | Weight |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Rows: [][]string{
				{"Alice", "20", "5'6", "120"},
				{"Bob", "30", "5'10", "150"},
				{"Charlie", "40", "6'0", "180"},
				{"David", "50", "6'2", "200"},
				{"Eve", "60", "6'4", "220"},
			},
			Align: map[int]string{
				0: "right",
			},
			Result: "|    Name | Age | Height | Weight |",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := NewTable(testCase.Header, testCase.Rows)
		table.Align = testCase.Align
		header := table.renderHeader()
		if header != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, header)
		}
	}

}

func TestRenderDivider(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Rows   [][]string
		Align  map[int]string
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Rows: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
			},
			Result: "| ----- | --- |",
		},
		{
			Header: []string{"Name", "Age"},
			Rows: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
			},
			Align: map[int]string{
				0: "center",
				1: "right",
			},
			Result: "| :---: | --: |",
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Rows: [][]string{
				{"Alice", "20", "5'6", "120"},
				{"Bob", "30", "5'10", "150"},
			},
			Align: map[int]string{
				0: "center",
				1: "right",
				2: "left",
			},
			Result: "| :---: | --: | :----- | ------ |",
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := NewTable(testCase.Header, testCase.Rows)
		table.Align = testCase.Align
		divider := table.renderDivider()
		if divider != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, divider)
		}
	}
}

func TestRenderRow(t *testing.T) {

	// Test Cases
	testCases := []struct {
		Header []string
		Rows   [][]string
		Align  map[int]string
		Result string
	}{
		{
			Header: []string{"Name", "Age"},
			Rows: [][]string{
				{"Alice", "20"},
				{"Bob", "30"},
			},
			Result: helpers.HereDoc(`
				| Alice | 20  |
				| Bob   | 30  |
			`),
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Rows: [][]string{
				{"Alice", "20", "5'6", "120"},
				{"Bob", "30", "5'10", "150"},
			},
			Result: helpers.HereDoc(`
				| Alice | 20  | 5'6    | 120    |
				| Bob   | 30  | 5'10   | 150    |
			`),
		},
		{
			Header: []string{"Name", "Age", "Height", "Weight"},
			Rows: [][]string{
				{"Alice", "20", "5'6", "120"},
				{"Bob", "30", "5'10", "150"},
				{"Charlie", "40", "6'0", "180"},
				{"David", "50", "6'2", "200"},
				{"Eve", "60", "6'4", "220"},
			},
			Align: map[int]string{
				0: "right",
			},
			Result: helpers.HereDoc(`
				|   Alice | 20  | 5'6    | 120    |
				|     Bob | 30  | 5'10   | 150    |
				| Charlie | 40  | 6'0    | 180    |
				|   David | 50  | 6'2    | 200    |
				|     Eve | 60  | 6'4    | 220    |
			`),
		},
	}

	// Run Test Cases
	for _, testCase := range testCases {
		table := NewTable(testCase.Header, testCase.Rows)
		table.Align = testCase.Align
		rows := table.renderRows()
		if rows != testCase.Result {
			t.Errorf("want:\n%v\n\ngot:\n%v", testCase.Result, rows)
		}
	}

}

func TestTable(t *testing.T) {

	// Create a table
	table := NewTable([]string{"Name", "Age"}, [][]string{
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
	table := NewTable([]string{"Name", "Age"}, [][]string{
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
