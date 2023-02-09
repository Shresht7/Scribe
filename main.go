package main

import (
	"fmt"
	"strings"
)

//* NODE INTERFACE *//

// Node is an interface for all nodes in the AST (Abstract Syntax Tree)
type Node interface {
	String() string
}

//* NODE LITERAL *//

// NodeLiteral is a node containing a literal value
type NodeLiteral struct {
	Value string
}

// Implement the Node interface for NodeLiteral
func (literal *NodeLiteral) String() string {
	return literal.Value
}

//* PARENT NODE *//

// ParentNode is a node containing other child nodes
type ParentNode struct {
	Nodes []Node
}

// Implement the Node interface for ParentNode.
// This is a recursive function that will call the String() method on all child nodes.
// The default implementation of the Node interface is very basic.
func (parent *ParentNode) String() string {
	res := []string{}
	for _, node := range parent.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, "")
}

// AppendChild appends the given nodes to the parent node
func (parent *ParentNode) AppendChild(nodes ...Node) *ParentNode {
	parent.Nodes = append(parent.Nodes, nodes...)
	return parent
}

//* NODE TEXT *//

// NodeText is a node containing text.
// It is a special type of ParentNode that joins the child nodes with a separator.
// The default separator is an empty string.
// The separator can be changed by calling the Join() method.
type NodeText struct {
	Separator string
	ParentNode
}

// Instantiate a new NodeText with the given contents
func Text(contents ...string) *NodeText {
	nodes := []Node{}
	for _, content := range contents {
		nodes = append(nodes, &NodeLiteral{Value: content})
	}
	return &NodeText{
		ParentNode: ParentNode{
			Nodes: nodes,
		},
	}
}

// Joins the text nodes with the given separator.
func (text *NodeText) Join(separator string) *NodeText {
	text.Separator = separator
	return text
}

// Implement the Node interface for NodeText
func (text *NodeText) String() string {
	res := []string{}
	for _, node := range text.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, text.Separator)
}

//* DOCUMENT *//

// Document is just a collection of text nodes
type Document struct {
	NodeText
}

// Instantiate a new document
func NewDocument(opts ...*Document) *Document {
	// Instantiate a new document with default options
	document := &Document{
		NodeText: NodeText{
			Separator: "\n",
		},
	}

	// If options are passed, override the defaults
	if len(opts) > 0 {
		options := opts[0]

		// Override default options with user options
		if options.Separator != "" {
			document.Separator = options.Separator
		}
	}

	// Return the document instance
	return document
}

//* MAIN FUNCTION *//

func main() {
	document := NewDocument()

	document.AppendChild(
		Text("Hello, world!"),
		Text("This", "is", "a", "test", "document.").Join("---"),
	)

	fmt.Println(document)
}
