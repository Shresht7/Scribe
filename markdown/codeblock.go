package markdown

//* CODEBLOCK *//

// A code block is a block of code that is surrounded
// by a line of three backticks (```) with optional language information.
type NodeCodeBlock struct {
	Languages []string
	NodeFencedBlock
}

// Create a new code block
func CodeBlock(contents string, languages ...string) *NodeCodeBlock {
	return &NodeCodeBlock{
		Languages: languages,
		NodeFencedBlock: NodeFencedBlock{
			Fence:    "```",
			Contents: contents,
		},
	}
}

// Implement the Node interface for code blocks
func (n *NodeCodeBlock) String() string {
	return n.WithMetadata(n.Languages...).String()
}
