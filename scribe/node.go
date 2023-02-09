package scribe

import "strings"

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

// ParentNode is a node containing other nodes
type ParentNode struct {
	Nodes []Node
}

// Implement the Node interface for ParentNode.
// This is a recursive function that will call the String() method on all child nodes.
// The default implementation of this Node interface is very basic.
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
