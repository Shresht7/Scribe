package markdown

import (
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

//* HEADING *//

// NodeHeading represents a heading in a markdown document.
type NodeHeading struct {
	Level int
	scribe.NodeText
}

// Instantiate a new heading with the given contents
func Heading(level int, contents ...any) *NodeHeading {
	// If the level is less than 1, set it to 1
	if level < 1 {
		level = 1
	}
	// If the level is greater than 6, set it to 6
	if level > 6 {
		level = 6
	}

	// Create a new heading
	heading := &NodeHeading{
		Level: level,
	}
	heading.WithSeparator(" ")

	// Append the contents to the heading
	heading.AppendChild(contents...)

	// Return the heading
	return heading
}

// Implement the Node interface for NodeHeading
func (heading *NodeHeading) String() string {
	return strings.Repeat("#", heading.Level) + " " + heading.NodeText.String()
}

// H1 is a convenience function for creating a heading of level 1.
func H1(contents ...any) *NodeHeading {
	return Heading(1, contents...)
}

// H2 is a convenience function for creating a heading of level 2.
func H2(contents ...any) *NodeHeading {
	return Heading(2, contents...)
}

// H3 is a convenience function for creating a heading of level 3.
func H3(contents ...any) *NodeHeading {
	return Heading(3, contents...)
}

// H4 is a convenience function for creating a heading of level 4.
func H4(contents ...any) *NodeHeading {
	return Heading(4, contents...)
}

// H5 is a convenience function for creating a heading of level 5.
func H5(contents ...any) *NodeHeading {
	return Heading(5, contents...)
}

// H6 is a convenience function for creating a heading of level 6.
func H6(contents ...any) *NodeHeading {
	return Heading(6, contents...)
}

// Write the heading to the document
func (doc *MarkdownDocument) WriteHeading(level int, contents string) *MarkdownDocument {
	doc.AppendChild(Heading(level, Text(contents)))
	return doc
}
