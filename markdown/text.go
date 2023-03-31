package markdown

import "github.com/Shresht7/Scribe/scribe"

// Alias for scribe.Node. This is done so that the markdown package can be used
// without importing scribe. The Node interface is used to represent any node
// in the AST (Abstract Syntax Tree). The only requirement is that the Node can
// be converted to the desired string using the String() method (i.e. it implements
// the Stringer interface).
type Node scribe.Node

// Instantiates a new NodeText with the given contents
func NewText(contents ...any) *scribe.NodeText {
	return scribe.NewText(contents...)
}

func Text(contents ...any) string {
	return NewText(contents...).String()
}
