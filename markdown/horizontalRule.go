package markdown

import "strings"

//* HORIZONTAL RULE *//

// NodeHorizontalRule represents a horizontal rule.
type NodeHorizontalRule struct {
	character rune
	Repeat    int
}

// Instantiate a new horizontal rule with the given style
func HorizontalRule(ch rune, rep int) *NodeHorizontalRule {
	// If the character is not a valid character, set it to '-'
	if ch != '*' && ch != '-' && ch != '_' {
		ch = '-'
	}

	// If the repeat is less than 3, set it to 3
	if rep < 3 {
		rep = 3
	}

	// Create a new horizontal rule
	return &NodeHorizontalRule{
		character: ch,
		Repeat:    rep,
	}
}

// Implement the Node interface for NodeHorizontalRule
func (hr *NodeHorizontalRule) String() string {
	return strings.Repeat(string(hr.character), hr.Repeat)
}

// Write a horizontal rule to the document
func (doc *MarkdownDocument) WriteHorizontalRule(ch rune, rep int) *MarkdownDocument {
	doc.AppendChild(HorizontalRule(ch, rep))
	return doc
}
