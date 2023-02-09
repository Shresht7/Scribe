package markdown

import (
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

//* HEADING *//

// NodeHeading represents a heading.
type NodeHeading struct {
	Level int
	scribe.NodeText
}

// Instantiate a new heading with the given contents
func Heading(level int, contents ...string) *NodeHeading {
	// If the level is less than 1, set it to 1
	if level < 1 {
		level = 1
	}
	// If the level is greater than 6, set it to 6
	if level > 6 {
		level = 6
	}

	// Return the heading
	return &NodeHeading{
		Level:    level,
		NodeText: *scribe.Text(contents...).Join(" "),
	}
}

// Implement the Node interface for NodeHeading
func (heading *NodeHeading) String() string {
	return strings.Repeat("#", heading.Level) + " " + heading.NodeText.String()
}
