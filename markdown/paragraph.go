package markdown

import "github.com/Shresht7/Scribe/scribe"

//* PARAGRAPH *//

// NodeParagraph represents a paragraph of text.
type NodeParagraph struct {
	scribe.NodeText
}

// Instantiate a new paragraph with the given contents
func Paragraph(contents ...scribe.Node) *NodeParagraph {
	return &NodeParagraph{
		scribe.NodeText{
			ParentNode: scribe.ParentNode{
				Nodes: contents,
			},
		},
	}
}

// Implement the Node interface for NodeParagraph
func (paragraph *NodeParagraph) String() string {
	return paragraph.NodeText.String()
}
