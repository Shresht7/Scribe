package scribe

import (
	"fmt"
	"io"
	"strings"
)

//* NODE INTERFACE *//

// Node is an interface that all nodes must satisfy.
type Node interface {
	String() string // Implement the Stringer interface
	io.Writer       // Implement the io.Writer interface
}

//* NODE LITERAL *//

// NodeLiteral is a node containing a literal value. It is the most
// basic implementation of the Node interface (Leaf Node).
type NodeLiteral struct {
	value string
}

// NewLiteral creates a new NodeLiteral with the given value.
func NewLiteral(v any) *NodeLiteral {
	value := fmt.Sprintf("%v", v)
	return &NodeLiteral{value}
}

// Implement the Node interface for NodeLiteral.
// Returns the literal string value of the node.
func (literal *NodeLiteral) String() string {
	return literal.value
}

// Implement the io.Writer interface for NodeLiteral.
// Writes the given bytes to the literal node.
func (literal *NodeLiteral) Write(p []byte) (int, error) {
	literal.value += string(p)
	return len(p), nil
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
func (container *NodeContainer) WithSeparator(separator string) *NodeContainer {
	container.Separator = separator
	return container
}

// Implement the Node interface for NodeContainer.
// This is a recursive function that will call the String() method on all child nodes.
// The default implementation of this Node interface is very basic.
func (container *NodeContainer) String() string {
	res := []string{}
	for _, node := range container.Nodes {
		res = append(res, node.String())
	}
	return strings.Join(res, container.Separator)
}

// AppendChild appends the given node to the NodeContainer.
func (container *NodeContainer) AppendChild(nodes ...any) *NodeContainer {

	// Iterate over the nodes
	for _, node := range nodes {
		// Append the node to the parent
		switch v := node.(type) {

		case Node:
			// If the node is a Node, then append it directly
			container.Nodes = append(container.Nodes, v)

		default:
			// If the node is anything else, then create a new NodeLiteral
			container.Nodes = append(container.Nodes, NewLiteral(v))

		}
	}

	// Return the parent
	return container
}

// Implement the io.Writer interface for NodeContainer.
// Writes the given bytes to the container node.
func (container *NodeContainer) Write(p []byte) (int, error) {
	// If the container is empty, then create a new NodeLiteral
	if len(container.Nodes) == 0 {
		container.Nodes = append(container.Nodes, &NodeLiteral{})
	}
	// Get the last node in the container
	lastNode := container.Nodes[len(container.Nodes)-1]
	// Write the bytes to the last node
	return lastNode.Write(p)
}
