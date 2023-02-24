package scribe

import "strings"

//* NODE INTERFACE *//

// Node is an interface for all nodes in the AST (Abstract Syntax Tree).
// The only requirement is that the Node can be converted to the desired string
// using the String() method (i.e. it implements the Stringer interface).
type Node interface {
	// String() returns the string representation of the node
	String() string
}

//* NODE LITERAL *//

// NodeLiteral is a node containing a literal value. It is the most
// basic implementation of the Node interface (Leaf Node).
type NodeLiteral struct {
	Value string
}

// Literal creates a new NodeLiteral with the given string
func Literal(value string) *NodeLiteral {
	return &NodeLiteral{Value: value}
}

// Implements the Node interface for NodeLiteral.
// Returns the literal string value of the node.
func (literal *NodeLiteral) String() string {
	return literal.Value
}

//* NODE PARENT *//

// NodeParent is a node containing other nodes. When rendered as a string,
// the child nodes are separated by the Separator string.
type NodeParent struct {
	Separator string
	Nodes     []Node
}

// Container creates a new NodeParent with the given nodes
func Container(nodes ...Node) *NodeParent {
	return &NodeParent{
		Separator: "",
		Nodes:     nodes,
	}
}

// WithSeparator sets the separator for the parent node
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
