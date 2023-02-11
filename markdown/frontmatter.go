package markdown

import "gopkg.in/yaml.v3"

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

// Create a new YAML FrontMatter block
func YAMLFrontMatter(contents any) (*NodeFrontMatter, error) {
	// Marshal the contents to YAML
	bytes, err := yaml.Marshal(contents)
	if err != nil {
		return nil, err
	}

	// Create the frontmatter block
	return FrontMatter("yaml", string(bytes)), nil

}
