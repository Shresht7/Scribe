package markdown

//* FRONTMATTER *//

// A frontmatter block is a fenced block with data
type NodeFrontMatter struct {
	*NodeFencedBlock
	Kind string
}

// Create a new frontmatter block
func FrontMatter(kind, contents string) *NodeFrontMatter {
	return &NodeFrontMatter{
		NodeFencedBlock: FencedBlock(contents),
		Kind:            kind,
	}
}
