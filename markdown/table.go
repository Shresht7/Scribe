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
	// The width of each column
	ColumnWidths []int
	// The alignment of each column
	Align map[int]string
}

// Create a new table node
func Table(header []string, rows [][]string) *NodeTable {
	// Determine the max width of each column
	columnWidths := helpers.DetermineWidths(append([][]string{header}, rows...))

	// Return the table
	return &NodeTable{
		Header:       header,
		Rows:         rows,
		ColumnWidths: columnWidths,
		Align:        map[int]string{},
	}
}

// Apply padding to a string based on the alignment
func (n *NodeTable) applyPadding(str string, col int) string {
	align := n.Align[col]
	width := n.ColumnWidths[col] - len(str)

	switch align {
	case "left":
		return helpers.PadRight(str, " ", width)
	case "right":
		return helpers.PadLeft(str, " ", width)
	case "center":
		return helpers.Pad(str, " ", width)
	default:
		return helpers.PadRight(str, " ", width)
	}
}

// Render the table header
func (n *NodeTable) renderHeader() string {
	res := []string{}
	for col, header := range n.Header {
		str := n.applyPadding(header, col)
		res = append(res, str)
	}
	return "| " + strings.Join(res, " | ") + " |"
}

// Render the table divider
func (n *NodeTable) renderDivider() string {
	res := []string{}
	for col := range n.Header {
		align := n.Align[col]
		width := n.ColumnWidths[col]
		switch align {
		case "left":
			res = append(res, ":"+strings.Repeat("-", width-1))
		case "right":
			res = append(res, strings.Repeat("-", width-1)+":")
		case "center":
			res = append(res, ":"+strings.Repeat("-", width-2)+":")
		default:
			res = append(res, strings.Repeat("-", width))
		}
	}
	return "| " + strings.Join(res, " | ") + " |"
}

// Render a table row
func (n *NodeTable) renderRow(row []string) string {
	res := []string{}
	for i, cell := range row {
		cell = n.applyPadding(cell, i)
		res = append(res, cell)
	}
	return "| " + strings.Join(res, " | ") + " |"
}

func (n *NodeTable) renderRows() string {
	res := []string{}
	for _, row := range n.Rows {
		res = append(res, n.renderRow(row))
	}
	return strings.Join(res, "\n")
}

// Render the table
func (n *NodeTable) String() string {
	res := []string{}

	// Render the header
	res = append(res, n.renderHeader())

	// Render the divider
	res = append(res, n.renderDivider())

	// Render the rows
	for _, row := range n.Rows {
		res = append(res, n.renderRow(row))
	}

	// Return the table as a string
	return strings.Join(res, "\n")
}

// Write Table to the document
func (doc *MarkdownDocument) WriteTable(header []string, rows [][]string) *MarkdownDocument {
	doc.AppendChild(Table(header, rows))
	return doc
}
