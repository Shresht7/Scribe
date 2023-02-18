package markdown

import "strings"

// * NodeTable *//

// NodeTable represents a markdown table
type NodeTable struct {
	// The table header
	Header []string
	// The table rows
	Rows [][]string
}

// Create a new table node
func Table(header []string, rows [][]string) *NodeTable {
	return &NodeTable{
		Header: header,
		Rows:   rows,
	}
}

func (n *NodeTable) String() string {
	res := []string{}
	// Create the header row
	res = append(res, "| "+strings.Join(n.Header, " | ")+" |")

	// Create the header divider
	dividers := []string{}
	for range n.Header {
		dividers = append(dividers, "---")
	}
	res = append(res, "| "+strings.Join(dividers, " | ")+" |")

	// Create the rows
	for _, row := range n.Rows {
		res = append(res, "| "+strings.Join(row, " | ")+" |")
	}

	// Return the table as a string
	return strings.Join(res, "\n")
}
