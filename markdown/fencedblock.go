package markdown

import "strings"

//* FENCED BLOCK *//

// A fenced code block is a block of text that is surrounded
// by a line of three or more hyphens (---) or tildes (~).
type NodeFencedBlock struct {
	Fence    string
	Contents string
	Metadata []string
}

// Create a new fenced block
func NewFencedBlock(contents ...string) *NodeFencedBlock {
	Contents := strings.Join(contents, "\n")
	return &NodeFencedBlock{
		Fence:    "---",
		Contents: Contents,
	}
}

// Set the fence style for the fenced block
func (n *NodeFencedBlock) WithFence(fence string) *NodeFencedBlock {
	n.Fence = fence
	return n
}

// Set the metadata for the fenced block
func (n *NodeFencedBlock) WithMetadata(metadata ...string) *NodeFencedBlock {
	n.Metadata = metadata
	return n
}

// Implement the Node interface for fenced blocks
func (n *NodeFencedBlock) String() string {
	res := []string{}

	// Create the top fence line and add metadata if it exists
	if len(n.Metadata) > 0 {
		res = append(res, n.Fence+strings.Join(n.Metadata, " "))
	} else {
		res = append(res, n.Fence)
	}

	// Add the contents and the bottom fence line
	res = append(res, n.Contents)
	res = append(res, n.Fence)

	// Return the fenced block as a string
	return strings.Join(res, "\n")
}

// Write a fenced block to the document
func (doc *MarkdownDocument) WriteFencedBlock(contents string) *MarkdownDocument {
	doc.AppendChild(NewFencedBlock(contents))
	return doc
}

// Create a new fenced block with the given contents
func FencedBlock(contents ...string) string {
	return NewFencedBlock(contents...).String()
}
