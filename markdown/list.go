package markdown

import (
	"strconv"
	"strings"

	"github.com/Shresht7/Scribe/scribe"
)

//* LIST *//

// List is a list of items.
type NodeList struct {
	Ordered   bool
	Decorator string
	Items     []scribe.Node
}

// Instantiate a new list with the given contents
func OrderedList(items []scribe.Node) *NodeList {
	return &NodeList{
		Ordered: true,
		Items:   items,
	}
}

// Instantiate a new list with the given contents
func UnorderedList(decorator string, items []scribe.Node) *NodeList {
	// Validate the decorator
	if decorator != "*" && decorator != "-" && decorator != "+" {
		decorator = "*"
	}

	// Return the list
	return &NodeList{
		Ordered:   false,
		Decorator: decorator,
		Items:     items,
	}
}

// Implement the Node interface for NodeList
func (list *NodeList) String() string {
	// Create a new string builder
	builder := strings.Builder{}

	// If the list is ordered, add the numbers
	if list.Ordered {
		// Iterate over the items
		for index, item := range list.Items {
			// Add the item to the builder
			builder.WriteString(strconv.Itoa(index+1) + ". " + item.String() + "\n")
		}
	} else {
		// Iterate over the items
		for _, item := range list.Items {
			// Add the item to the builder
			builder.WriteString(list.Decorator + " " + item.String() + "\n")
		}
	}

	// Return the string
	return builder.String()
}
