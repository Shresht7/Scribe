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
func NewParagraph(contents ...any) *NodeParagraph {
	paragraph := &NodeParagraph{}
	paragraph.WithSeparator("\n")
	paragraph.AppendChild(contents...)
	return paragraph
}

// Write a paragraph to the document
func (doc *MarkdownDocument) WriteParagraph(contents ...string) *MarkdownDocument {
	paragraph := NewParagraph()
	for _, content := range contents {
		paragraph.AppendChild(NewText(content))
	}
	doc.AppendChild(paragraph)
	return doc
}

func Paragraph(contents ...any) string {
	return NewParagraph(contents...).String()
}
