package markdown

import (
	"strings"

	"github.com/Shresht7/Scribe/helpers"
)

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

// Render the table header
func (n *NodeTable) renderHeader(widths []int) string {
	res := []string{}
	for i, header := range n.Header {
		res = append(res, header+strings.Repeat(" ", widths[i]-len(header)))
	}
	return "| " + strings.Join(res, " | ") + " |"
}

// Render the table divider
func (n *NodeTable) renderDivider(widths []int) string {
	res := []string{}
	for i := range n.Header {
		res = append(res, strings.Repeat("-", widths[i]))
	}
	return "| " + strings.Join(res, " | ") + " |"
}

// Render a table row
func (n *NodeTable) renderRow(row []string, widths []int) string {
	res := []string{}
	for i, cell := range row {
		res = append(res, cell+strings.Repeat(" ", widths[i]-len(cell)))
	}
	return "| " + strings.Join(res, " | ") + " |"
}

// Render the table
func (n *NodeTable) String() string {
	res := []string{}

	// Determine the max width of each column
	table := append([][]string{n.Header}, n.Rows...)
	widths := helpers.DetermineWidths(table)

	// Render the header
	res = append(res, n.renderHeader(widths))

	// Render the divider
	res = append(res, n.renderDivider(widths))

	// Render the rows
	for _, row := range n.Rows {
		res = append(res, n.renderRow(row, widths))
	}

	// Return the table as a string
	return strings.Join(res, "\n")
}
