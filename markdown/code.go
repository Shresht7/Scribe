package markdown

import "github.com/Shresht7/Scribe/scribe"

//* CODE *//

// Code represents inline code text.
type NodeCode struct {
	scribe.NodeText
}

// Instantiate a new code text with the given contents
func Code(contents ...scribe.Node) *NodeCode {
	// Create a new code
	code := &NodeCode{}
	code.WithSeparator(" ")

	// Append the contents to the code
	code.AppendChild(contents...)

	// Return the code
	return code
}

// Implement the Node interface for NodeCode
func (code *NodeCode) String() string {
	return "`" + code.NodeText.String() + "`"
}
