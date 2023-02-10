package markdown

import (
	"github.com/Shresht7/Scribe/scribe"
)

//* PARAGRAPH *//

// NodeParagraph represents a paragraph of text.
type NodeParagraph struct {
	scribe.NodeText
}

// Instantiate a new paragraph with the given contents
func Paragraph(contents ...scribe.Node) *NodeParagraph {
	paragraph := &NodeParagraph{}
	paragraph.WithSeparator(" ")
	paragraph.AppendChild(contents...)
	return paragraph
}
