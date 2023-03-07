package markdown

//* CODEBLOCK *//

// A code block is a block of code that is surrounded
// by a line of three backticks (```) with optional language information.
type NodeCodeBlock struct {
	Languages []string
	NodeFencedBlock
}

// Create a new code block
func NewCodeBlock(contents string, languages ...string) *NodeCodeBlock {
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

// Write a code block to the document
func (doc *MarkdownDocument) WriteCodeBlock(contents string, languages ...string) *MarkdownDocument {
	doc.AppendChild(NewCodeBlock(contents, languages...))
	return doc
}

// Create a new code block with the given contents
func CodeBlock(contents string, languages ...string) string {
	return NewCodeBlock(contents, languages...).String()
}
