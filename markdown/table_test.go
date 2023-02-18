package markdown

import (
	"testing"

	"github.com/Shresht7/Scribe/helpers"
)

func TestTable(t *testing.T) {

	// Create a table
	table := Table([]string{"Name", "Age"}, [][]string{
		{"Alice", "20"},
		{"Bob", "30"},
	})

	// Check the table string
	if table.String() != helpers.HereDoc(`
		| Name | Age |
		| --- | --- |
		| Alice | 20 |
		| Bob | 30 |
	`) {
		t.Error("Table string is not correct")
	}

}
