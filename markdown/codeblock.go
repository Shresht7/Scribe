package markdown

import "strings"

//* CODEBLOCK *//

// NodeCodeBlock represents a code block.
type NodeCodeBlock struct {
	languages []string
	code      string
}

// Instantiate a new code block with the given contents
func CodeBlock(code string, languages ...string) *NodeCodeBlock {
	// Create a new code block
	codeblock := &NodeCodeBlock{
		languages: languages,
		code:      code,
	}

	// Return the code block
	return codeblock
}

// Implement the Node interface for NodeCodeBlock
func (codeblock *NodeCodeBlock) String() string {
	// Create the language string
	languages := ""
	if len(codeblock.languages) > 0 {
		languages = strings.Join(codeblock.languages, " ")
	}

	// Create the code block string
	return "```" + languages + "\n" + codeblock.code + "\n```"
}
