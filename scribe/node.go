package scribe

import "strings"

//* NODE INTERFACE *//

// Node is an interface for all nodes in the AST (Abstract Syntax Tree)
type Node interface {
	// String() returns the string representation of the node
	String() string
}

//* NODE LITERAL *//

// NodeLiteral is a node containing a literal value
type NodeLiteral struct {
	Value string
}

// Implement the Node interface for NodeLiteral.
// Returns the literal string value of the node.
func (literal *NodeLiteral) String() string {
	return literal.Value
}

//* NODE PARENT *//

// NodeParent is a node containing other nodes
type NodeParent struct {
	Separator string
	Nodes     []Node
}

// WithSeparator sets the separator for the parent node.
func (parent *NodeParent) WithSeparator(separator string) *NodeParent {
	parent.Separator = separator
	return parent
}

// Implement the Node interface for ParentNode.
// This is a recursive function that will call the String() method on all child nodes.
// The default implementation of this Node interface is very basic.
func (parent *NodeParent) String() string {
	res := []string{}
	for _, node := range parent.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, parent.Separator)
}

// AppendChild appends the given node to the parent node.
func (parent *NodeParent) AppendChild(nodes ...Node) *NodeParent {
	parent.Nodes = append(parent.Nodes, nodes...)
	return parent
}
