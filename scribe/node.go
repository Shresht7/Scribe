package scribe

import (
	"strings"
)

//* NODE INTERFACE *//

// Node is an interface for all nodes in the AST (Abstract Syntax Tree).
// The only requirement is that the Node can be converted to the desired string
// using the String() method (i.e. it implements the Stringer interface).
type Node interface {
	// ? Consider adding a `Type()` method to return the type of the node
	// String() returns the string representation of the node
	String() string
}

//* NODE LITERAL *//

// NodeLiteral is a node containing a literal value. It is the most
// basic implementation of the Node interface (Leaf Node).
type NodeLiteral struct {
	// TODO: Accept anything and try to convert it to a string
	value string
}

// NewLiteral creates a new NodeLiteral with the given value.
func NewLiteral(value string) *NodeLiteral {
	return &NodeLiteral{value}
}

// Implement the Node interface for NodeLiteral.
// Returns the literal string value of the node.
func (literal *NodeLiteral) String() string {
	return literal.value
}

//* NODE CONTAINER *//

// NodeContainer is a node containing other nodes. When rendered as a string,
// the child nodes are separated by the `Separator` string.
type NodeContainer struct {
	Separator string
	Nodes     []Node
}

// NewContainer creates a new NodeContainer with the given nodes. The
// nodes can be of type string or implement the Node interface. If the node
// is a string, then it is converted to a NodeLiteral. If the node is a Node,
// then it is appended directly.
func NewContainer(nodes ...Node) *NodeContainer {
	return &NodeContainer{
		Separator: "",
		Nodes:     nodes,
	}
}

// WithSeparator sets the separator for the NodeContainer
func (n *NodeContainer) WithSeparator(separator string) *NodeContainer {
	n.Separator = separator
	return n
}

// Implement the Node interface for NodeContainer.
// This is a recursive function that will call the String() method on all child nodes.
// The default implementation of this Node interface is very basic.
func (n *NodeContainer) String() string {
	res := []string{}
	for _, node := range n.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, n.Separator)
}

// AppendChild appends the given node to the NodeContainer.
func (n *NodeContainer) AppendChild(nodes ...any) *NodeContainer {

	// Iterate over the nodes
	for _, node := range nodes {
		// Append the node to the parent
		switch v := node.(type) {

		case string:
			// If the node is a string, then create a new NodeLiteral
			n.Nodes = append(n.Nodes, &NodeLiteral{v})
			// TODO: Change this case to default and let NodeLiteral handle the conversion

		case Node:
			// If the node is a Node, then append it directly
			n.Nodes = append(n.Nodes, v)

		}
	}

	// Return the parent
	return n
}
