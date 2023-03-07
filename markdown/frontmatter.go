package markdown

//* FRONTMATTER *//

// A frontmatter block is a fenced block with data
type NodeFrontMatter struct {
	*NodeFencedBlock
	Kind string
}

// Create a new frontmatter block
func NewFrontMatter(kind, contents string) *NodeFrontMatter {
	return &NodeFrontMatter{
		NodeFencedBlock: NewFencedBlock(contents),
		Kind:            kind,
	}
}

// Write a frontmatter block to the document
func (doc *MarkdownDocument) WriteFrontMatter(kind, contents string) *MarkdownDocument {
	doc.AppendChild(NewFrontMatter(kind, contents))
	return doc
}

// Create a new frontmatter block with the given contents
func FrontMatter(kind, contents string) string {
	return NewFrontMatter(kind, contents).String()
}
